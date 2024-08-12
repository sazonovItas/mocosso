package logger

import (
	"context"
	"io"
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	stdout = "stdout"
	stderr = "stderr"
)

// NewSlogInterceptor function returns LogFunc from slog logger
func NewSlogInterceptor(logger *slog.Logger) LogFunc {
	return func(lvl Level, msg string, keysAndValues ...any) {
		logger.Log(context.Background(), SlogLevel(lvl), msg, keysAndValues...)
	}
}

// SlogLevel function converts levels of slog logger
func SlogLevel(lvl Level) slog.Level {
	switch lvl {
	case DebugLevel:
		return slog.LevelDebug
	case InfoLevel:
		return slog.LevelInfo
	case WarnLevel:
		return slog.LevelWarn
	case ErrorLevel:
		return slog.LevelError
	case PanicLevel:
		return slog.LevelError
	case FatalLevel:
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

// NewZapInterceptor function returns LogFunc from zap logger
func NewZapInterceptor(logger *zap.Logger) LogFunc {
	return func(lvl Level, msg string, keysAndValues ...any) {
		logger.Sugar().Logw(ZapLevel(lvl), msg, keysAndValues...)
	}
}

// ZapLevel function converts levels of zap logger
func ZapLevel(lvl Level) zapcore.Level {
	switch lvl {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

// NewWriter function returns new writer from string that could be
// stdout, stderr or file's name to log at. If file does not exist create new.
func NewWriter(out string) (wrt io.WriteCloser, err error) {
	switch out {
	case stdout:
		return os.Stdout, nil
	case stderr:
		return os.Stderr, nil
	default:
		// checks file is exists
		_, err = os.Stat(out)
		if err != nil {
			return nil, err
		}

		// create a new file if does not exists or open to append
		wrt, err = os.OpenFile(out, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
		if err != nil {
			return nil, err
		}

		return wrt, err
	}
}

// NewWriter function returns new writer from string that could be
// stdout, stderr or file's name to log at. If file does not exist create new.
// If any or errors occurred it panics.
func MustNewWriter(out string) io.WriteCloser {
	switch out {
	case stdout:
		return os.Stdout
	case stderr:
		return os.Stderr
	default:
		// checks file is exists
		_, err := os.Stat(out)
		if err != nil {
			panic(err)
		}

		// create a new file if does not exists or open to append
		wrt, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
		if err != nil {
			panic(err)
		}

		return wrt
	}
}
