package handlers

import (
	"github.com/hansel-app/hansel/internal/adapters/primary/handlers/internal"
	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories"
	"github.com/hansel-app/hansel/protobuf/gemsapi"
	"google.golang.org/grpc"
)

// RegisterServices registers all gRPC services available with the gRPC server.
func RegisterServices(s *grpc.Server) {
	gemsapi.RegisterGemServiceServer(s, internal.NewGemService(&repositories.GemRepository{}))
}
