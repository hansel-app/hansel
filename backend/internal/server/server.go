package server

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/internal/adapters/primary/handlers"
)

func New(db *sqlx.DB) *grpc.Server {
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(authenticateRequest)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authenticateRequest)),
	)
	handlers.RegisterServices(s, db)
	return s
}

func authenticateRequest(ctx context.Context) (context.Context, error) {
	// TODO: Implement authentication logic
	return ctx, nil
}
