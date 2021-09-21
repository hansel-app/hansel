package handlers

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/internal/adapters/primary/handlers/internal"
	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories"
	"github.com/hansel-app/hansel/internal/auth"
	"github.com/hansel-app/hansel/protobuf/authapi"
	"github.com/hansel-app/hansel/protobuf/friendsapi"
	"github.com/hansel-app/hansel/protobuf/gemsapi"
	"github.com/hansel-app/hansel/protobuf/usersapi"
)

// RegisterServices registers all gRPC services available with the gRPC server.
func RegisterServices(s *grpc.Server, db *sqlx.DB, jwtManager *auth.JWTManager) {
	userRepository := repositories.NewUserRepository(db)
	userService := internal.NewUserService(userRepository)
	authapi.RegisterAuthServiceServer(s, internal.NewAuthService(userRepository, jwtManager))
	gemsapi.RegisterGemServiceServer(s, internal.NewGemService(repositories.NewGemRepository(db)))
	usersapi.RegisterUserServiceServer(s, userService)
	friendsapi.RegisterFriendServiceServer(s, userService)
}
