package internal

import (
	"context"
	"fmt"
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

func (h *gemService) Drop(c context.Context, r *gemsapi.DropRequest) (*gemsapi.DropResponse, error) {
	gem, err := h.usecases.Drop(r.Message)
	if err != nil {
		return nil, err
	}

	return &gemsapi.DropResponse{
		Id:      gem.ID,
		Message: gem.Message,
	}, nil
}

func (h *gemService) SayHello(c context.Context, r *gemsapi.SayHelloRequest) (*gemsapi.SayHelloResponse, error) {
	return &gemsapi.SayHelloResponse{
		Message: fmt.Sprintf("Hello, %s", r.Name),
	}, nil
}
