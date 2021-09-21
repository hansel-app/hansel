package internal

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hansel-app/hansel/internal/contextkeys"
	"github.com/hansel-app/hansel/internal/core/domain/gems"
	"github.com/hansel-app/hansel/protobuf/gemsapi"
)

type gemService struct {
	gemsapi.UnimplementedGemServiceServer
	usecases gems.UseCases
}

func NewGemService(repository gems.Repository) *gemService {
	return &gemService{
		usecases: *gems.NewUseCases(repository),
	}
}

func (s *gemService) Drop(c context.Context, r *gemsapi.DropRequest) (*gemsapi.DropResponse, error) {
	gem, err := s.usecases.Drop(r.Message)
	if err != nil {
		return nil, err
	}

	return &gemsapi.DropResponse{
		Id:      gem.ID,
		Message: gem.Message,
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
	r *gemsapi.GetPendingCollectionForUserRequest,
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
