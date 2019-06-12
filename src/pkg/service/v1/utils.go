package v1

import (
	"google.golang.org/api/option"
	"context"
	"google.golang.org/api/oauth2/v2"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/c-beel/userman/src/models"
	"fmt"
	"github.com/c-beel/userman/src/configman"
)

func NewUsermanServer(config configman.Config) (*UsermanServer, error) {
	db, err := gorm.Open("sqlite3", config.DBAddress)
	if err != nil {
		return nil, err
	}
	return &UsermanServer{
		DB:                db,
		GoogleOAuthAPIKey: config.GoogleOAuthAPIKey,
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

func (server *UsermanServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
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
