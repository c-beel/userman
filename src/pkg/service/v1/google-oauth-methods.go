package v1

import (
	"context"
	"github.com/c-beel/userman/src/pkg/api/v1"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/c-beel/userman/src/models"
	"github.com/jinzhu/gorm"
)

func (server UsermanServer) GetOrCreateUserByIdToken(ctx context.Context, req *v1.GetOrCreateUserByIdTokenRequest) (*v1.GetOrCreateUserByIdTokenResponse, error) {
	var email string
	email, err := server.getEmailByIdToken(req.IdToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to get id token email with error %v", err)
	}
	var user models.User
	userFound := true
	if err := server.DB.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			userFound = false
		} else {
			return nil, status.Errorf(codes.NotFound, "Could not find any user with email(%s) with error %v", email, err)
		}
	}
	if userFound {
		var responseUser v1.User
		userModelsToGrpc(&user, &responseUser)
		return &v1.GetOrCreateUserByIdTokenResponse{
			User:    &responseUser,
			NewUser: false,
		}, nil
	}
	user = models.User{
		Email: email,
	}
	if err := server.DB.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to create user with email(%s) with error %v", email, err)
	}
	var responseUser v1.User
	userModelsToGrpc(&user, &responseUser)
	return &v1.GetOrCreateUserByIdTokenResponse{
		User:    &responseUser,
		NewUser: true,
	}, nil
}
