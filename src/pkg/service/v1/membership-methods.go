package v1

import (
	"../../api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"../../../models"
	"context"
)

func (server *UsermanServer) AddUserToGroup(ctx context.Context, req *v1.AddUserToGroupRequest) (*v1.AddUserToGroupResponse, error) {
	var user models.User
	if err := server.DB.First(&user, req.Uid).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get user with this id(%d) with error %v", req.Uid, err)
	}
	for _, reqGroup := range req.Groups {
		var group models.Group
		if err := server.DB.Where(&models.Group{Name: reqGroup.Name}).First(&group).Error; err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to get group with name(%s) with error %v", reqGroup.Name, err)
		}
		membership := models.Membership{
			UID: user.ID,
			GID: group.ID,
		}
		if err := server.DB.Create(&membership).Error; err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to add membership of group(%s) with error %v", group.Name, err)
		}
	}
	return &v1.AddUserToGroupResponse{}, nil
}

func (server *UsermanServer) RemoveUserFromGroup(ctx context.Context, req *v1.RemoveUserFromGroupRequest) (*v1.RemoveUserFromGroupResponse, error) {
	var user models.User
	if err := server.DB.First(&user, req.Uid).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get user with this id(%d) with error %v", req.Uid, err)
	}
	for _, reqGroup := range req.Groups {
		var group models.Group
		if err := server.DB.Where(&models.Group{Name: reqGroup.Name}).First(&group).Error; err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to get group with name(%s) with error %v", reqGroup.Name, err)
		}
		membership := models.Membership{
			UID: user.ID,
			GID: group.ID,
		}
		if err := server.DB.Where(&membership).Delete(&models.Membership{}).Error; err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to remove membership in group(%s) with error %v", group.Name, err)
		}
	}
	return &v1.RemoveUserFromGroupResponse{}, nil
}

func (server *UsermanServer) SetUserGroups(ctx context.Context, req *v1.SetUserGroupsRequest) (*v1.SetUserGroupsResponse, error) {
	var user models.User
	if err := server.DB.First(&user, req.Uid).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get user with this id(%d) with error %v", req.Uid, err)
	}
	if err := server.DB.Where(&models.Membership{UID: user.ID}).Delete(&models.Membership{}).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to delete memberships of user(%d) with error %v", user.ID, err)
	}
	for _, reqGroup := range req.Groups {
		var group models.Group
		if err := server.DB.Where(&models.Group{Name: reqGroup.Name}).First(&group).Error; err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to get group with name(%s) with error %v", reqGroup.Name, err)
		}
		membership := models.Membership{
			UID: user.ID,
			GID: group.ID,
		}
		if err := server.DB.Create(&membership).Error; err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to add membership of group(%s) with error %v", group.Name, err)
		}
	}
	return &v1.SetUserGroupsResponse{}, nil
}

func (server *UsermanServer) IsMemberOf(ctx context.Context, req *v1.IsMemberOfRequest) (*v1.IsMemberOfResponse, error) {
	var user models.User
	if err := server.DB.First(&user, req.Uid).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get user with this id(%d) with error %v", req.Uid, err)
	}
	server.DB.Where(&models.Membership{UID: user.ID}).Delete(&models.Membership{})
	var group models.Group
	if err := server.DB.Where(&models.Group{Name: req.Group.Name}).First(&group).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to get group with name(%s) with error %v", req.Group.Name, err)
	}

	var memberships []models.Membership
	if err := server.DB.Where(&models.Membership{UID: user.ID, GID: group.ID}).Find(&memberships).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to get memberships of user(%d) with error %v", user.ID, err)
	}
	return &v1.IsMemberOfResponse{
		Yes: len(memberships) > 0,
	}, nil
}
