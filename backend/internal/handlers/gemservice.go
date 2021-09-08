package handlers

import (
	"context"
	"fmt"
	"github.com/hansel-app/hansel/protobuf/gem"
)

type gemService struct {
	gem.UnimplementedGemServiceServer
}

func (h *gemService) Drop(c context.Context, r *gem.DropRequest) (*gem.DropResponse, error) {
	// TODO: Replace placeholder return value.
	return &gem.DropResponse{
		Id:      1,
		Message: r.Message,
	}, nil
}

func (h *gemService) HelloWorld(c context.Context, r *gem.HelloWorldRequest) (*gem.HelloWorldResponse, error) {
	return &gem.HelloWorldResponse{
		Message: fmt.Sprintf("Hello, %s", r.Name),
	}, nil
}
