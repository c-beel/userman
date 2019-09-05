package v1

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	// apiVersion is version of API provided by server
	apiVersion = "v1"
)

type UsermanServer struct {
	DB                *gorm.DB
	GoogleOAuthAPIKey string
}

/*
// Create new User
func (server *UsermanServer) UpdateUserOld(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	// check if the API version requested by client is supported by server
	var err error
	if err = server.checkAPI(req.Api); err != nil {
		return nil, err
	}

	var user models.User
	user = models.User{
		User: *req.User,
	}
	user.Id = 0
	user.Email, err = server.getEmailByIdToken(req.IdToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "failed to authorize user-> "+err.Error())
	}

	var dbUsers []models.User
	if err = server.DB.Where(&models.User{User: v1.User{Email: user.Email}}).Find(&dbUsers).Error; err != nil {
		return nil, status.Error(codes.Unknown, "failed to get users with this email.")
	}
	if len(dbUsers) == 0 {
		// insert User entity data
		if err = server.DB.Create(&user).Error; err != nil {
			return nil, status.Error(codes.Unknown, "failed to add user-> "+err.Error())
		}
		return &v1.UpdateUserResponse{
			Api: apiVersion,
			Id:  user.Id,
		}, nil
	} else {
		// update User entity data
		fmt.Println(user.FirstName)
		fmt.Println(dbUsers[0].FirstName)
		if err = server.DB.Model(&dbUsers[0]).Where("id = ?", dbUsers[0].Id).Update(&user).Error; err != nil {
			return nil, status.Error(codes.Unknown, "failed to update user-> "+err.Error())
		}
		return &v1.UpdateUserResponse{
			Api: apiVersion,
			Id:  dbUsers[0].Id,
		}, nil
	}

}

// get a User
func (server *UsermanServer) GetUserById(ctx context.Context, req *v1.GetUserByIdRequest) (*v1.GetUserByIdResponse, error) {
	// check if the API version requested by client is supported by server
	if err := server.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if !server.isAuthenticated(req.IdToken) {
		return nil, status.Error(codes.Unauthenticated, "failed to authorize id token")
	}

	var user models.User
	if err := server.DB.Where(&models.User{User: v1.User{Id: req.Id}}).First(&user).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed with error %v", err.Error())
	}

	return &v1.GetUserByIdResponse{
		Api:  apiVersion,
		User: &user.User,
	}, nil
}
*/
