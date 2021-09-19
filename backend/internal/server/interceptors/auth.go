package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/hansel-app/hansel/internal/auth"
	"github.com/hansel-app/hansel/internal/contextkeys"
	"github.com/hansel-app/hansel/internal/server/serverstream"
)

type AuthInterceptor struct {
	jwtManager *auth.JWTManager
}

func NewAuthInterceptor(jwtManager *auth.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager: jwtManager,
	}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx, err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (i *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx, err := i.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		ss := serverstream.NewServerStreamWithContext(stream)
		ss.SetContext(ctx)

		return handler(srv, stream)
	}
}

func (i *AuthInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	// TODO: Add exceptions for login and registration methods.
	log.Println(method)

	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	values := m["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "missing authorization token")
	}

	accessToken := values[0]
	claims, err := i.jwtManager.Verify(accessToken)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "invalid access token: %v", err)
	}

	// Add the user ID to the context for use in requests.
	ctx = context.WithValue(ctx, contextkeys.UserID, claims.UserID)
	return ctx, nil
}
