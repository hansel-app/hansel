package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/internal/adapters/primary/handlers"
	"github.com/hansel-app/hansel/internal/auth"
	"github.com/hansel-app/hansel/internal/server/interceptors"
)

func New(db *sqlx.DB, jwtManager *auth.JWTManager) *grpc.Server {
	authInterceptor := interceptors.NewAuthInterceptor(jwtManager)
	loggingInterceptor := interceptors.NewLoggingInterceptor()

	// Note that the interceptors are executed in the order in which they are declared here.
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			loggingInterceptor.Unary(),
			authInterceptor.Unary(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			loggingInterceptor.Stream(),
			authInterceptor.Stream(),
		)),
	)

	handlers.RegisterServices(s, db)
	return s
}
