package logging

import (
	"context"
	"log/slog"
)

// https://dev.to/ilyakaznacheev/where-to-place-logger-in-golang-13o3

type ctxLogger struct{}

// ContextWithLogger adds logger to context.
func ContextWithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context.
func LoggerFromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*slog.Logger); ok {
		return l
	}

	panic(ErrLoggerNotFound)
}
