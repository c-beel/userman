package v1

import (
	"google.golang.org/api/option"
	"context"
	"google.golang.org/api/oauth2/v2"
	"github.com/jinzhu/gorm"
	"github.com/c-beel/userman/src/models"
	"github.com/c-beel/userman/src/configman"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/validations"
	"errors"
)

func NewUsermanServer(cfg *configman.MainConfig) (*UsermanServer, error) {
	dbUri := cfg.DBConfig.GetDBUri()
	db, err := gorm.Open(cfg.DBConfig.Type, dbUri)
	validations.RegisterCallbacks(db)
	if err != nil {
		return nil, err
	}
	return &UsermanServer{
		DB:           db,
		GOAuthConfig: cfg.GOAuthConfig,
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
	oauth2Service, err := oauth2.NewService(ctx, option.WithAPIKey(server.GOAuthConfig.APIKey))
	if err != nil {
		return "", err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return "", err
	}
	if tokenInfo.Audience != server.GOAuthConfig.Issuer {
		return "", errors.New("invalid audience")
	}
	if !tokenInfo.VerifiedEmail {
		return "", errors.New("non verified email")
	}
	return tokenInfo.Email, nil
}
