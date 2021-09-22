package internal

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hansel-app/hansel/internal/auth"
	"github.com/hansel-app/hansel/internal/core/domain/users"
	"github.com/hansel-app/hansel/protobuf/authapi"
)

type authService struct {
	authapi.UnimplementedAuthServiceServer
	usecases   users.UseCases
	jwtManager auth.JWTManager
}

func NewAuthService(repository users.Repository, jwtManager *auth.JWTManager) *authService {
	return &authService{
		usecases:   *users.NewUseCases(repository),
		jwtManager: *jwtManager,
	}
}

func (s *authService) Login(ctx context.Context, r *authapi.LoginRequest) (*authapi.LoginResponse, error) {
	userID, err := s.usecases.AuthenticatePassword(r.Username, r.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, err := s.jwtManager.Generate(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to generate JWT")
	}

	return &authapi.LoginResponse{
		AccessToken: token,
	}, nil
}

func (s *authService) Register(
	c context.Context,
	r *authapi.RegisterRequest,
) (*authapi.RegisterResponse, error) {
	userId, err := s.usecases.Register(r.DisplayName, r.Username, r.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authapi.RegisterResponse{
		UserId: userId,
	}, nil
}
