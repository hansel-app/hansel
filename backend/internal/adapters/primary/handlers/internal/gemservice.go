package internal

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hansel-app/hansel/internal/contextkeys"
	"github.com/hansel-app/hansel/internal/core/domain/gems"
	"github.com/hansel-app/hansel/internal/core/domain/users"
	"github.com/hansel-app/hansel/protobuf/gemsapi"
	"github.com/hansel-app/hansel/protobuf/usersapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type gemService struct {
	gemsapi.UnimplementedGemServiceServer
	usecases      gems.UseCases
	usersUsecases users.UseCases
}

func NewGemService(repository gems.Repository, userRepository users.Repository) *gemService {
	return &gemService{
		usecases:      *gems.NewUseCases(repository),
		usersUsecases: *users.NewUseCases(userRepository),
	}
}

func (s *gemService) Drop(c context.Context, r *gemsapi.DropRequest) (*gemsapi.DropResponse, error) {
	userID, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	gemMessage := r.GetGemMessage()
	gem := gems.Gem{
		Message:    gemMessage.GetMessage(),
		Latitude:   gemMessage.GetLatitude(),
		Longitude:  gemMessage.GetLongitude(),
		CreatorId:  userID,
		ReceiverId: gemMessage.GetReceiverId(),
		ReceivedAt: nil,
		Color:      gems.GemColor(gemMessage.GetColor()),
		Attachment: gemMessage.GetAttachment(),
	}
	droppedGemId, err := s.usecases.Drop(&gem)
	if err != nil {
		return nil, err
	}

	return &gemsapi.DropResponse{
		DroppedGemId: droppedGemId,
	}, nil
}

func (s *gemService) Get(c context.Context, r *gemsapi.GetRequest) (*gemsapi.GetResponse, error) {
	gem, err := s.usecases.Get(r.Id)
	if err != nil {
		return nil, err
	}

	return &gemsapi.GetResponse{
		Id:      gem.ID,
		Message: gem.Message,
	}, nil
}

func (s *gemService) GetPendingCollectionForUser(
	c context.Context,
	_ *empty.Empty,
) (*gemsapi.GetPendingCollectionForUserResponse, error) {
	userID, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	gems, err := s.usecases.GetPendingCollectionByUser(userID)
	if err != nil {
		return nil, err
	}
	if len(gems) == 0 {
		return &gemsapi.GetPendingCollectionForUserResponse{
			Gems: nil,
		}, nil
	}

	// Store all user ids to be retrieved as keys of the map
	// Map is used primarily as a set here
	userIdsSet := make(map[int64]bool)
	for _, gem := range gems {
		userIdsSet[gem.CreatorId] = true
		userIdsSet[gem.ReceiverId] = true
	}

	// Retrieve all userIds of the map
	userIds := make([]int64, len(userIdsSet))
	i := 0
	for k := range userIdsSet {
		userIds[i] = k
		i++
	}

	// retrieve all users based on keys
	userInfoMap, err := s.usersUsecases.GetUsersByIds(userIds)
	if err != nil {
		return nil, err
	}

	processedGems := []*gemsapi.Gem{}

	for _, gem := range gems {
		creator := userInfoMap[gem.CreatorId]
		receiver := userInfoMap[gem.ReceiverId]
		processedGem := gemsapi.Gem{
			Id:        gem.ID,
			Message:   gem.Message,
			Latitude:  gem.Latitude,
			Longitude: gem.Longitude,
			Creator: &usersapi.User{
				UserId:      creator.ID,
				DisplayName: creator.DisplayName,
				Username:    creator.Username,
				Avatar:      creator.Avatar,
			},
			CreatedAt: timestamppb.New(gem.CreatedAt),
			Receiver: &usersapi.User{
				UserId:      receiver.ID,
				DisplayName: receiver.DisplayName,
				Username:    receiver.Username,
				Avatar:      receiver.Avatar,
			},
			ReceivedAt: nil,
			Color:      gemsapi.GemColor(gem.Color),
		}
		processedGems = append(processedGems, &processedGem)
	}

	return &gemsapi.GetPendingCollectionForUserResponse{
		Gems: processedGems,
	}, nil
}

func (s *gemService) PickUp(
	c context.Context,
	r *gemsapi.PickUpRequest,
) (*emptypb.Empty, error) {
	err := s.usecases.PickUpGem(r.GetId())
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (s *gemService) GetGemLogs(
	c context.Context,
	_ *empty.Empty,
) (*gemsapi.GemLogs, error) {
	userID, ok := c.Value(contextkeys.UserID).(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "unable to retrieve user ID from context")
	}

	friendIdToGemLogsMap, err := s.usecases.GetGemLogs(userID)
	if err != nil {
		return nil, err
	}
	// Retrieve list of relevant user ids
	userIds := make([]int64, len(friendIdToGemLogsMap))
	i := 0
	for friendId := range friendIdToGemLogsMap {
		userIds[i] = friendId
		i++
	}
	// Retrieve user information about self as well
	userIds = append(userIds, userID)

	userInfoMap, err := s.usersUsecases.GetUsersByIds(userIds)
	if err != nil {
		return nil, err
	}

	gemLogs := make(map[int64]*gemsapi.GemLogsWithFriend)
	for friendId, gems := range friendIdToGemLogsMap {
		// TODO: maybe abstract out this method of processing gems
		// and maybe abstract out the mappers...
		processedGems := []*gemsapi.Gem{}

		for _, gem := range gems {
			creator := userInfoMap[gem.CreatorId]
			receiver := userInfoMap[gem.ReceiverId]

			var receivedAt *timestamppb.Timestamp
			if gem.ReceivedAt != nil {
				receivedAt = timestamppb.New(*gem.ReceivedAt)
			}

			processedGem := gemsapi.Gem{
				Id:        gem.ID,
				Message:   gem.Message,
				Latitude:  gem.Latitude,
				Longitude: gem.Longitude,
				Creator: &usersapi.User{
					UserId:      creator.ID,
					DisplayName: creator.DisplayName,
					Username:    creator.Username,
					Avatar:      creator.Avatar,
				},
				CreatedAt: timestamppb.New(gem.CreatedAt),
				Receiver: &usersapi.User{
					UserId:      receiver.ID,
					DisplayName: receiver.DisplayName,
					Username:    receiver.Username,
					Avatar:      receiver.Avatar,
				},
				ReceivedAt: receivedAt,
				Color:      gemsapi.GemColor(gem.Color),
			}
			processedGems = append(processedGems, &processedGem)
		}

		friend := userInfoMap[friendId]
		gemLogsWithFriend := gemsapi.GemLogsWithFriend{
			Gems: processedGems,
			Friend: &usersapi.User{
				UserId:      friend.ID,
				DisplayName: friend.DisplayName,
				Username:    friend.Username,
				Avatar:      friend.Avatar,
			},
		}
		gemLogs[friendId] = &gemLogsWithFriend
	}

	return &gemsapi.GemLogs{
		GemLogs: gemLogs,
	}, nil
}
