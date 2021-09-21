package internal

import (
	"context"

	"github.com/hansel-app/hansel/internal/core/domain/users"
	"github.com/hansel-app/hansel/protobuf/friendsapi"
)

type userService struct {
	friendsapi.UnimplementedFriendServiceServer
	usecases users.UseCases
}

func NewUserService(repository users.Repository) *userService {
	return &userService{
		usecases: *users.NewUseCases(repository),
	}
}

func (s *userService) GetFriends(c context.Context, r *friendsapi.GetFriendsRequest) (*friendsapi.GetFriendsResponse, error) {
	friends, err := s.usecases.GetFriends(r.UserId)
	if err != nil {
		return nil, err
	}

	var friendsInfo []*friendsapi.FriendInfo
	for _, f := range friends {
		friendsInfo = append(friendsInfo, 
		&friendsapi.FriendInfo{
			UserId: f.ID,
			DisplayName: f.DisplayName,
			Username: f.Username,
		})
	}
	return &friendsapi.GetFriendsResponse{
		Friends: friendsInfo,
	}, nil
}


func (s *userService) GetFriendRequests(c context.Context, r *friendsapi.GetPendingFriendRequestsRequest) (*friendsapi.GetPendingFriendRequestsResponse, error) {
	friends, err := s.usecases.GetFriendRequests(r.UserId)
	if err != nil {
		return nil, err
	}

	var friendsInfo []*friendsapi.FriendInfo
	for _, f := range friends {
		friendsInfo = append(friendsInfo, 
		&friendsapi.FriendInfo{
			UserId: f.ID,
			DisplayName: f.DisplayName,
			Username: f.Username,
		})
	}
	return &friendsapi.GetPendingFriendRequestsResponse{
		Friends: friendsInfo,
	}, nil
}


