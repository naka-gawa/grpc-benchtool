package logging

import (
	"context"
	"log/slog"
)

type ctxKey struct{}

var loggerKey = ctxKey{}

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}
