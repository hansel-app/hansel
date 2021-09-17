package internal

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

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
	gem, err := s.usecases.Drop(r.gem)
	if err != nil {
		return nil, err
	}

	return &gemsapi.DropResponse{
		droppedGem: r.gem
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

func (s *gemService) GetPendingCollectionByUser(
	c context.Context,
	r *gemsapi.GetPendingCollectionByUserRequest,
) (*gemsapi.GetPendingCollectionByUserResponse, error) {
	gems, err := s.usecases.GetPendingCollectionByUser(r.UserId)
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
		}
		processedGems = append(processedGems, &processedGem)
	}

	return &gemsapi.GetPendingCollectionByUserResponse{
		Gems: processedGems,
	}, nil
}

func (s *gemService) SayHello(c context.Context, r *gemsapi.SayHelloRequest) (*gemsapi.SayHelloResponse, error) {
	return &gemsapi.SayHelloResponse{
		Message: fmt.Sprintf("Hello, %s", r.Name),
	}, nil
}
