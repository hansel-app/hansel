package internal

import (
	"context"

	"github.com/hansel-app/hansel/internal/core/domain/friends"
	"github.com/hansel-app/hansel/protobuf/friendsapi"
)

type friendService struct {
	friendsapi.UnimplementedFriendServiceServer
	usecases friends.UseCases
}

func NewFriendService(repository friends.Repository) *friendService {
	return &friendService{
		usecases: *friends.NewUseCases(repository),
	}
}

func (s *friendService) GetFriends(c context.Context, r *friendsapi.GetFriendsRequest) (*friendsapi.GetFriendsResponse, error) {
	friends, err := s.usecases.Get(r.UserId)
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

