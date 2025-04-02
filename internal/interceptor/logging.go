package interceptor

import (
	"context"
	"log/slog"
	"time"

	"github.com/naka-gawa/grpc-benchtool/internal/logging"
	codepb "google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func withRPCLogging(ctx context.Context, method string, reqID string, handler func(context.Context) error) error {
	logger := logging.FromContext(ctx).With(
		slog.String("method", method),
		slog.String("request_id", reqID),
	)

	logger.InfoContext(ctx, "RPC started")

	start := time.Now()
	err := handler(ctx)

	st, _ := status.FromError(err)
	msg := st.Message()

	code := status.Code(err)
	logger.InfoContext(ctx, "RPC finished",
		slog.Int("status_code", int(codepb.Code(code))),
		slog.String("status_name", code.String()),
		slog.String("message", msg),
		slog.Duration("duration", time.Since(start)),
		slog.Any("error", err),
	)

	return err
}

func UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		reqID, _ := ctx.Value(RequestIDKey).(string)

		var resp interface{}
		err := withRPCLogging(ctx, info.FullMethod, reqID, func(ctx context.Context) error {
			var innerErr error
			resp, innerErr = handler(ctx, req)
			return innerErr
		})
		return resp, err
	}
}

func StreamLoggingInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		reqID, _ := ss.Context().Value(RequestIDKey).(string)
		return withRPCLogging(ss.Context(), info.FullMethod, reqID, func(ctx context.Context) error {
			return handler(srv, ss)
		})
	}
}
