package internal

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hansel-app/hansel/internal/contextkeys"
	"github.com/hansel-app/hansel/internal/core/domain/users"
	"github.com/hansel-app/hansel/protobuf/friendsapi"
	"github.com/hansel-app/hansel/protobuf/usersapi"
)

type userService struct {
	usersapi.UnimplementedUserServiceServer
	friendsapi.UnimplementedFriendServiceServer
	usecases users.UseCases
}

func NewUserService(repository users.Repository) *userService {
	return &userService{
		usecases: *users.NewUseCases(repository),
	}
}

func (s *userService) GetProfile(
	c context.Context,
	r *usersapi.ProfileRequest,
) (*usersapi.ProfileResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	userInfo, err := s.usecases.Get(userId)
	if err != nil {
		return nil, err
	}

	getFriendsResponse, err := s.GetFriends(c, &usersapi.GetFriendsRequest{})
	if err != nil {
		return nil, err
	}

	return &usersapi.ProfileResponse{
		Info: &usersapi.PersonInfo{
			UserId:      userInfo.ID,
			DisplayName: userInfo.DisplayName,
			Username:    userInfo.Username,
			Avatar:      userInfo.Avatar,
		},
		Friends: getFriendsResponse.GetFriends(),
	}, nil
}

func (s *userService) GetFriends(
	c context.Context,
	r *usersapi.GetFriendsRequest,
) (*usersapi.GetFriendsResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	friends, err := s.usecases.GetFriends(userId)
	if err != nil {
		return nil, err
	}

	var friendsInfo []*usersapi.PersonInfo
	for _, f := range friends {
		friendsInfo = append(friendsInfo,
			&usersapi.PersonInfo{
				UserId:      f.ID,
				DisplayName: f.DisplayName,
				Username:    f.Username,
				Avatar:      f.Avatar,
			})
	}

	return &usersapi.GetFriendsResponse{
		Friends: friendsInfo,
	}, nil
}

func (s *userService) SearchByUsername(
	c context.Context,
	r *usersapi.SearchByUsernameRequest,
) (*usersapi.SearchByUsernameResponse, error) {
	matchingUsers, err := s.usecases.SearchByUsername(r.SearchQuery)
	if err != nil {
		return nil, err
	}

	matchingUsersInfo := make([]*usersapi.PersonInfo, len(matchingUsers))
	for i, user := range matchingUsers {
		matchingUsersInfo[i] = &usersapi.PersonInfo{
			UserId:      user.ID,
			DisplayName: user.DisplayName,
			Username:    user.Username,
			Avatar:      user.Avatar,
		}
	}

	return &usersapi.SearchByUsernameResponse{
		Users: matchingUsersInfo,
	}, nil
}

func (s *userService) GetFriendRequests(
	c context.Context,
	r *usersapi.GetPendingFriendRequestsRequest,
) (*usersapi.GetPendingFriendRequestsResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	requests, err := s.usecases.GetFriendRequests(userId)
	if err != nil {
		return nil, err
	}

	var friendRequests []*usersapi.PendingFriendRequest
	for _, f := range requests {
		requester := f.Requester
		friendRequests = append(friendRequests,
			&usersapi.PendingFriendRequest{
				Requester: &usersapi.PersonInfo{
					UserId:      requester.ID,
					DisplayName: requester.DisplayName,
					Username:    requester.Username,
					Avatar:      requester.Avatar,
				},
				SentAt: timestamppb.New(f.RequestedAt),
			},
		)
	}

	return &usersapi.GetPendingFriendRequestsResponse{
		FriendRequests: friendRequests,
	}, nil
}

func (s *userService) AddFriendRequest(
	c context.Context,
	r *friendsapi.FriendRequest,
) (*emptypb.Empty, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	err := s.usecases.AddFriendRequest(userId, r.ReceiverId)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (s *userService) AcceptFriendRequest(
	c context.Context,
	r *friendsapi.FriendRequest,
) (*emptypb.Empty, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return new(emptypb.Empty), status.Error(
			codes.Internal, "unable to retrieve user ID from context",
		)
	}

	err := s.usecases.AcceptFriendRequest(r.RequesterId, userId)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (s *userService) DeclineFriendRequest(
	c context.Context,
	r *friendsapi.FriendRequest,
) (*emptypb.Empty, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	err := s.usecases.DeclineFriendRequest(r.RequesterId, userId)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}
