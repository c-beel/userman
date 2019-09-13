package v1

import (
	"google.golang.org/grpc/codes"
	"context"
	"google.golang.org/grpc/status"
	"github.com/c-beel/userman/src/pkg/api/v1"
	"github.com/c-beel/userman/src/models"
)

func (server UsermanServer) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	var user models.User

	userGrpcToModels(req.User, &user)

	if err := server.DB.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to create user with error %v.", err)
	}
	var newUser v1.User
	userModelsToGrpc(&user, &newUser)
	return &v1.CreateUserResponse{
		User: &newUser,
	}, nil
}

func (server UsermanServer) ReadUser(ctx context.Context, req *v1.ReadUserRequest) (*v1.ReadUserResponse, error) {
	var user models.User
	if err := server.DB.First(&user, req.Username).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get user with this username(%s) with error %v", req.Username, err)
	}
	var newUser v1.User
	userModelsToGrpc(&user, &newUser)
	return &v1.ReadUserResponse{
		User: &newUser,
	}, nil
}

func (server UsermanServer) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	var user models.User

	if err := server.DB.First(&user, req.User.Username).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get user with this username(%s) with error %v", req.User.Username, err)
	}

	userGrpcToModels(req.User, &user)

	if err := server.DB.Save(&user).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to save user updates with error %v.", err)
	}
	var updatedUser v1.User
	userModelsToGrpc(&user, &updatedUser)
	return &v1.UpdateUserResponse{
		User: &updatedUser,
	}, nil
}

func (server UsermanServer) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	var user models.User
	if err := server.DB.First(&user, req.Username).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get user with this username(%s) with error %v", req.Username, err)
	}
	if err := server.DB.Delete(&user).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to delete user with this user(%s) with error %v", req.Username, err)
	}
	var deletedUser v1.User
	userModelsToGrpc(&user, &deletedUser)
	return &v1.DeleteUserResponse{
		User: &deletedUser,
	}, nil
}
