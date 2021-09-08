package handlers

import (
	"context"
	"fmt"
	"github.com/hansel-app/hansel/protobuf/gems"
)

type gemService struct {
	gems.UnimplementedGemServiceServer
}

func (h *gemService) Drop(c context.Context, r *gems.DropRequest) (*gems.DropResponse, error) {
	// TODO: Replace placeholder return value.
	return &gems.DropResponse{
		Id:      1,
		Message: r.Message,
	}, nil
}

func (h *gemService) SayHello(c context.Context, r *gems.SayHelloRequest) (*gems.SayHelloResponse, error) {
	return &gems.SayHelloResponse{
		Message: fmt.Sprintf("Hello, %s", r.Name),
	}, nil
}
