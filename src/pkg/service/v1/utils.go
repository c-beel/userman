package v1

import (
	"google.golang.org/api/option"
	"context"
	"google.golang.org/api/oauth2/v2"
	"github.com/jinzhu/gorm"
	"github.com/c-beel/userman/src/models"
	"fmt"
	"github.com/c-beel/userman/src/configman"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewUsermanServer(cfg *configman.MainConfig) (*UsermanServer, error) {
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DBConfig.Address, cfg.DBConfig.Port, cfg.DBConfig.Username, cfg.DBConfig.Database, cfg.DBConfig.Password)
	db, err := gorm.Open(cfg.DBConfig.Type, dbUri)
	if err != nil {
		return nil, err
	}
	return &UsermanServer{
		DB:                db,
		GoogleOAuthAPIKey: cfg.GoogleOAuthAPIKey,
	}, nil
}

func (server *UsermanServer) AutoMigrate() (err error) {
	if err = server.DB.AutoMigrate(&models.User{}).Error; err != nil {
		return err
	}
	if err = server.DB.AutoMigrate(&models.Group{}).Error; err != nil {
		return err
	}
	if err = server.DB.AutoMigrate(&models.Membership{}).Error; err != nil {
		return err
	}
	return nil
}

func (server *UsermanServer) getEmailByIdToken(idToken string) (string, error) {
	ctx := context.Background()
	oauth2Service, err := oauth2.NewService(ctx, option.WithAPIKey(server.GoogleOAuthAPIKey))
	if err != nil {
		return "", err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return "", err
	}
	return tokenInfo.Email, nil
}

func (server *UsermanServer) isAuthenticated(idToken string) bool {
	ctx := context.Background()
	oauth2Service, err := oauth2.NewService(ctx, option.WithAPIKey(server.GoogleOAuthAPIKey))
	if err != nil {
		return false
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return false
	}
	fmt.Println(tokenInfo.Email, tokenInfo.VerifiedEmail)
	return tokenInfo.VerifiedEmail
}
