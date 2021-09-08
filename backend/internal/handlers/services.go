package handlers

import (
	"github.com/hansel-app/hansel/protobuf/gem"
	"google.golang.org/grpc"
)

// RegisterServices registers all gRPC services available with the gRPC server.
func RegisterServices(s *grpc.Server) {
	gem.RegisterGemServiceServer(s, &gemService{})
}
