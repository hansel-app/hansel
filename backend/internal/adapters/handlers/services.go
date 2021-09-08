package handlers

import (
	"github.com/hansel-app/hansel/internal/adapters/handlers/internal"
	"github.com/hansel-app/hansel/internal/adapters/repositories"
	"github.com/hansel-app/hansel/protobuf/gemsapi"
	"google.golang.org/grpc"
)

// RegisterServices registers all gRPC services available with the gRPC server.
func RegisterServices(s *grpc.Server) {
	gemsapi.RegisterGemServiceServer(s, internal.NewGemService(&repositories.GemRepository{}))
}
