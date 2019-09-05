package v1

import (
	"github.com/c-beel/userman/src/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"context"
	"github.com/c-beel/userman/src/models"
	"google.golang.org/grpc/status"
)

func (server *UsermanServer) CreateGroup(ctx context.Context, req *v1.CreateGroupRequest) (*v1.CreateGroupResponse, error) {
	var group models.Group

	groupGrpcToModels(req.Group, &group)

	if err := server.DB.Create(&group).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to create group with error %v.", err)
	}
	var newGroup v1.Group
	groupModelsToGrpc(&group, &newGroup)
	return &v1.CreateGroupResponse{
		Group: &newGroup,
	}, nil
}

func (server *UsermanServer) ReadGroupList(ctx context.Context, req *v1.ReadGroupListRequest) (*v1.ReadGroupListResponse, error) {
	var groupList []models.Group
	if err := server.DB.Find(&groupList).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get groups with error %v", err)
	}
	responseGroups := make([]*v1.Group, len(groupList))
	for index, group := range groupList {
		var appendingGroup v1.Group
		groupModelsToGrpc(&group, &appendingGroup)
		responseGroups[index] = &appendingGroup
	}
	return &v1.ReadGroupListResponse{
		Groups: responseGroups,
	}, nil
}

func (server *UsermanServer) DeleteGroup(ctx context.Context, req *v1.DeleteGroupRequest) (*v1.DeleteGroupResponse, error) {
	var group models.Group
	if err := server.DB.Where(&models.Group{Name: req.Group.Name}).First(&group).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get group with this name(%s) with error %v", req.Group.Name, err)
	}
	if err := server.DB.Delete(&group).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to get group with this name(%s) with error %v", req.Group.Name, err)
	}
	var deletedGroup v1.Group
	groupModelsToGrpc(&group, &deletedGroup)
	return &v1.DeleteGroupResponse{
		Group: &deletedGroup,
	}, nil
}
