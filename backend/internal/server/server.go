package server

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/internal/adapters/primary/handlers"
	"github.com/hansel-app/hansel/internal/auth"
	"github.com/hansel-app/hansel/internal/server/interceptors"
)

func New(db *sqlx.DB, jwtManager *auth.JWTManager) *grpc.Server {
	authInterceptor := interceptors.NewAuthInterceptor(jwtManager)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
		grpc.StreamInterceptor(authInterceptor.Stream()),
	)
	handlers.RegisterServices(s, db)
	return s
}
