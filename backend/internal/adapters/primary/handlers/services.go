package handlers

import (
	"github.com/hansel-app/hansel/internal/adapters/primary/handlers/internal"
	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories"
	"github.com/hansel-app/hansel/internal/auth"
	"github.com/hansel-app/hansel/protobuf/authapi"
	"github.com/hansel-app/hansel/protobuf/gemsapi"
	"github.com/hansel-app/hansel/protobuf/usersapi"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

// RegisterServices registers all gRPC services available with the gRPC server.
func RegisterServices(s *grpc.Server, db *sqlx.DB, jwtManager *auth.JWTManager) {
	gemRepository := repositories.NewGemRepository(db)
	userRepository := repositories.NewUserRepository(db)

	authService := internal.NewAuthService(userRepository, jwtManager)
	gemService := internal.NewGemService(gemRepository, userRepository)
	userService := internal.NewUserService(userRepository)

	authapi.RegisterAuthServiceServer(s, authService)
	gemsapi.RegisterGemServiceServer(s, gemService)
	usersapi.RegisterUserServiceServer(s, userService)
}
