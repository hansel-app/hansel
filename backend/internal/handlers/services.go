package handlers

import (
	"github.com/hansel-app/hansel/protobuf/gems"
	"google.golang.org/grpc"
)

// RegisterServices registers all gRPC services available with the gRPC server.
func RegisterServices(s *grpc.Server) {
	gems.RegisterGemServiceServer(s, &gemService{})
}
