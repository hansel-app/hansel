package handlers

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/internal/adapters/primary/handlers/internal"
	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories"
	"github.com/hansel-app/hansel/protobuf/gemsapi"
)

// RegisterServices registers all gRPC services available with the gRPC server.
func RegisterServices(s *grpc.Server, db *sqlx.DB) {
	gemsapi.RegisterGemServiceServer(s, internal.NewGemService(repositories.NewGemRepository(db)))
}
