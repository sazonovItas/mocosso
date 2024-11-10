package logger

import (
	"context"

	"go.uber.org/zap"
)

func init() {
	_ = ConfigureLogger()
}

func Sync() error {
	return getLogger().Sync()
}

func NamedContext(ctx context.Context, name string) *zap.Logger {
	return FromContext(ctx).Named(name)
}

func WithContext(ctx context.Context, fields ...zap.Field) *zap.Logger {
	return FromContext(ctx).With(fields...)
}

func DebugContext(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Debug(msg, fields...)
}

func InfoContext(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Info(msg, fields...)
}

func WarnContext(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Warn(msg, fields...)
}

func ErrorContext(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Error(msg, fields...)
}

func FatalContext(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Fatal(msg, fields...)
}

func PanicContetxt(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Panic(msg, fields...)
}

func SugarContext(ctx context.Context) *zap.SugaredLogger {
	return FromContext(ctx).Sugar()
}

func getLogger() *zap.Logger {
	return CreateLogger()
}
