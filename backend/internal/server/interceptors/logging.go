package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type LoggingInterceptor struct{}

func NewLoggingInterceptor() *LoggingInterceptor {
	return &LoggingInterceptor{}
}

func (i *LoggingInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		i.Log(info.FullMethod)
		return handler(ctx, req)
	}
}

func (i *LoggingInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		i.Log(info.FullMethod)
		return handler(srv, stream)
	}
}

func (i *LoggingInterceptor) Log(method string) {
	log.Println(method)
}
