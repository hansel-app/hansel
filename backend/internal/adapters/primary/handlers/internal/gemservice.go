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

	processedGems := []*gemsapi.Gem{}

	for _, gem := range gems {
		processedGem := gemsapi.Gem{
			Id:         gem.ID,
			Message:    gem.Message,
			Latitude:   gem.Latitude,
			Longitude:  gem.Longitude,
			CreatorId:  gem.CreatorId,
			CreatedAt:  timestamppb.New(gem.CreatedAt),
			ReceiverId: gem.ReceiverId,
			Color:      gemsapi.GemColor(gem.Color),
		}
		processedGems = append(processedGems, &processedGem)
	}

	return &gemsapi.GetPendingCollectionForUserResponse{
		Gems: processedGems,
	}, nil
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
	for id := range friendIdToGemLogsMap {
		userIds[i] = id
		i++
	}

	userInfoMap, err := s.usersUsecases.GetUsersByIds(userIds)
	if err != nil {
		return nil, err
	}

	gemLogs := make(map[int64]*gemsapi.GemLogsWithFriend)
	for id, gems := range friendIdToGemLogsMap {
		f := userInfoMap[id]

		// TODO: maybe abstract out this method of processing gems
		// and maybe abstract out the mappers...
		processedGems := []*gemsapi.Gem{}
		for _, gem := range gems {
			processedGem := gemsapi.Gem{
				Id:         gem.ID,
				Message:    gem.Message,
				Latitude:   gem.Latitude,
				Longitude:  gem.Longitude,
				CreatorId:  gem.CreatorId,
				CreatedAt:  timestamppb.New(gem.CreatedAt),
				ReceiverId: gem.ReceiverId,
				Color:      gemsapi.GemColor(gem.Color),
			}
			processedGems = append(processedGems, &processedGem)
		}

		gemLogsWithFriend := gemsapi.GemLogsWithFriend{
			Friend: &usersapi.User{
				UserId:      f.ID,
				DisplayName: f.DisplayName,
				Username:    f.Username,
				Avatar:      f.Avatar,
			},
			Gems: processedGems,
		}
		gemLogs[id] = &gemLogsWithFriend
	}

	return &gemsapi.GemLogs{
		GemLogs: gemLogs,
	}, nil
}
