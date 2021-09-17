package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/internal/adapters/primary/handlers"
	"github.com/hansel-app/hansel/internal/auth"
)

func New(db *sqlx.DB) *grpc.Server {
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(auth.AuthenticateRequest)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(auth.AuthenticateRequest)),
	)
	handlers.RegisterServices(s, db)
	return s
}
