package serverstream

import (
	"context"

	"google.golang.org/grpc"
)

// ServerStreamWithContext is a grpc.ServerStream that allows its context to be modified.
// It does so by overwriting the Context() function of grpc.ServerStream.
type ServerStreamWithContext struct {
	grpc.ServerStream
	ctx context.Context
}

func NewServerStreamWithContext(ss grpc.ServerStream) ServerStreamWithContext {
	return ServerStreamWithContext{
		ServerStream: ss,
		ctx:          ss.Context(),
	}
}

func (ss *ServerStreamWithContext) Context() context.Context {
	return ss.ctx
}

func (ss *ServerStreamWithContext) SetContext(ctx context.Context) {
	ss.ctx = ctx
}
