package internal

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hansel-app/hansel/internal/contextkeys"
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

func (s *userService) GetFriends(
	c context.Context,
	r *friendsapi.GetFriendsRequest,
) (*friendsapi.GetFriendsResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	friends, err := s.usecases.GetFriends(userId)
	if err != nil {
		return nil, err
	}

	var friendsInfo []*friendsapi.FriendInfo
	for _, f := range friends {
		friendsInfo = append(friendsInfo,
			&friendsapi.FriendInfo{
				UserId:      f.ID,
				DisplayName: f.DisplayName,
				Username:    f.Username,
			})
	}
	return &friendsapi.GetFriendsResponse{
		Friends: friendsInfo,
	}, nil
}

func (s *userService) GetFriendRequests(
	c context.Context,
	r *friendsapi.GetPendingFriendRequestsRequest,
) (*friendsapi.GetPendingFriendRequestsResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	friends, err := s.usecases.GetFriendRequests(userId)
	if err != nil {
		return nil, err
	}

	var friendRequests []*friendsapi.PendingFriendRequest
	for _, f := range friends {
		friendRequests = append(friendRequests,
			&friendsapi.PendingFriendRequest{
				Requester: &friendsapi.FriendInfo{
				UserId:      f.ID,
				DisplayName: f.DisplayName,
				Username:    f.Username,
				},
			},
		)
	}

	return &friendsapi.GetPendingFriendRequestsResponse{
		FriendRequests: friendRequests,
	}, nil
}

func (s *userService) AddFriendRequest(
	c context.Context,
	r *friendsapi.FriendRequest,
) error {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	err := s.usecases.AddFriendRequest(userId, r.ReceiverId)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) AcceptFriendRequest(
	c context.Context,
	r *friendsapi.FriendRequest,
) error {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	err := s.usecases.AcceptFriendRequest(r.RequesterId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) DeclineFriendRequest(
	c context.Context,
	r *friendsapi.FriendRequest,
) error {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	err := s.usecases.DeclineFriendRequest(r.RequesterId, userId)
	if err != nil {
		return err
	}
	return nil
}
