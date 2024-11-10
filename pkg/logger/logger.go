package logger

import (
	"context"

	"go.uber.org/zap"
)

type loggerCtxKey struct{}

func ToContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerCtxKey{}).(*zap.Logger); ok {
		return logger
	}

	return getLogger()
}
