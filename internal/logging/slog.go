package logging

import (
	"context"
	"log/slog"
	"os"
)

type Logger interface {
	InfoContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
	With(args ...any) Logger
}

type slogLogger struct {
	logger *slog.Logger
}

func (l *slogLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *slogLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}

func (l *slogLogger) With(args ...any) Logger {
	return &slogLogger{logger: l.logger.With(args...)}
}

func Init() Logger {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(l)
	return &slogLogger{logger: l}
}
