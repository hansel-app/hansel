package internal

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hansel-app/hansel/internal/contextkeys"
	"github.com/hansel-app/hansel/internal/core/domain/users"
	"github.com/hansel-app/hansel/protobuf/usersapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userService struct {
	usersapi.UnimplementedUserServiceServer
	usecases users.UseCases
}

func NewUserService(repository users.Repository) *userService {
	return &userService{
		usecases: *users.NewUseCases(repository),
	}
}

func (s *userService) GetOwnProfile(
	c context.Context,
	_ *empty.Empty,
) (*usersapi.GetOwnProfileResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	userInfo, err := s.usecases.Get(userId)
	if err != nil {
		return nil, err
	}

	getFriendsResponse, err := s.GetFriends(c, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return &usersapi.GetOwnProfileResponse{
		Info: &usersapi.User{
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
	_ *empty.Empty,
) (*usersapi.GetFriendsResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	friends, err := s.usecases.GetFriends(userId)
	if err != nil {
		return nil, err
	}

	var friendsInfo []*usersapi.User
	for _, f := range friends {
		friendsInfo = append(friendsInfo,
			&usersapi.User{
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

func (s *userService) GetPendingFriends(
	c context.Context,
	_ *empty.Empty,
) (*usersapi.GetPendingFriendsResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	pendingFriends, err := s.usecases.GetPendingFriends(userId)
	if err != nil {
		return nil, err
	}

	var friendsInfo []*usersapi.User
	for _, p := range pendingFriends {
		friendsInfo = append(friendsInfo,
			&usersapi.User{
				UserId:      p.ID,
				DisplayName: p.DisplayName,
				Username:    p.Username,
				Avatar:      p.Avatar,
			})
	}

	return &usersapi.GetPendingFriendsResponse{
		PendingFriends: friendsInfo,
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

	matchingUsersInfo := make([]*usersapi.User, len(matchingUsers))
	for i, user := range matchingUsers {
		matchingUsersInfo[i] = &usersapi.User{
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
	_ *empty.Empty,
) (*usersapi.GetFriendRequestsResponse, error) {
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
				Requester: &usersapi.User{
					UserId:      requester.ID,
					DisplayName: requester.DisplayName,
					Username:    requester.Username,
					Avatar:      requester.Avatar,
				},
				CreatedAt: timestamppb.New(f.CreatedAt),
			},
		)
	}

	return &usersapi.GetFriendRequestsResponse{
		FriendRequests: friendRequests,
	}, nil
}

func (s *userService) AddFriendRequest(
	c context.Context,
	r *usersapi.FriendRequest,
) (*usersapi.AddFriendRequestResponse, error) {
	userId, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	addFriendRequestStatus, err := s.usecases.AddFriendRequest(userId, r.ReceiverId)
	if err != nil {
		return nil, err
	}

	return &usersapi.AddFriendRequestResponse{
		Status: usersapi.AddFriendRequestStatus(addFriendRequestStatus),
	}, nil
}

func (s *userService) AcceptFriendRequest(
	c context.Context,
	r *usersapi.FriendRequest,
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
	r *usersapi.FriendRequest,
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

func (s *userService) EditOwnProfile(
	c context.Context,
	r *usersapi.EditProfileRequest,
) (*emptypb.Empty, error) {
	userID, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	if r.GetNewAvatar() != nil {
		err := s.usecases.UpdateAvatar(userID, r.GetNewAvatar())
		if err != nil {
			return nil, err
		}
	}

	if r.GetNewDisplayName() != "" {
		err := s.usecases.UpdateDisplayName(userID, r.GetNewDisplayName())
		if err != nil {
			return nil, err
		}
	}

	return new(emptypb.Empty), nil
}
