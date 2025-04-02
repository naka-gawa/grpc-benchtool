package interceptor

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ctxKey string

const RequestIDKey ctxKey = "request-id"

func UnaryRequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		reqID := uuid.NewString()
		ctx = context.WithValue(ctx, RequestIDKey, reqID)
		return handler(ctx, req)
	}
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context { return w.ctx }

func StreamRequestIDInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		reqID := uuid.NewString()
		ctx := context.WithValue(ss.Context(), RequestIDKey, reqID)
		return handler(srv, &wrappedStream{ServerStream: ss, ctx: ctx})
	}
}
