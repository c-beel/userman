package v1

import (
	"github.com/c-beel/userman/src/pkg/api/v1"
	"github.com/c-beel/userman/src/models"
)

func userGrpcToModels(grpcUser *v1.User, modelUser *models.User, includeId bool) {
	if includeId {
		modelUser.ID = uint(grpcUser.Id)
	}
	modelUser.Username = grpcUser.Username
	modelUser.Nickname = grpcUser.Nickname
	modelUser.Email = grpcUser.Email
	modelUser.FirstName = grpcUser.FirstName
	modelUser.LastName = grpcUser.LastName
}

func userModelsToGrpc(modelUser *models.User, grpcUser *v1.User, includeId bool) {
	if includeId {
		grpcUser.Id = int64(modelUser.ID)
	}
	grpcUser.Username = modelUser.Username
	grpcUser.Nickname = modelUser.Nickname
	grpcUser.Email = modelUser.Email
	grpcUser.FirstName = modelUser.FirstName
	grpcUser.LastName = modelUser.LastName
}

func groupGrpcToModels(grpcGroup *v1.Group, modelGroup *models.Group) {
	modelGroup.Name = grpcGroup.Name
}

func groupModelsToGrpc(modelGroup *models.Group, grpcGroup *v1.Group) {
	grpcGroup.Name = modelGroup.Name
}

func groupListModelsToGrpc(modelGroups *[]models.Group, grpcGroups *[]*v1.Group) {
	*grpcGroups = make([]*v1.Group, len(*modelGroups))
	for index, group := range *modelGroups {
		var appendingGroup v1.Group
		groupModelsToGrpc(&group, &appendingGroup)
		(*grpcGroups)[index] = &appendingGroup
	}
}

func groupListGrpcToModels(grpcGroups *[]*v1.Group, modelGroups *[]models.Group) {
	*modelGroups = make([]models.Group, len(*grpcGroups))
	for index, group := range *grpcGroups {
		var appendingGroup models.Group
		groupGrpcToModels(group, &appendingGroup)
		(*modelGroups)[index] = appendingGroup
	}
}
