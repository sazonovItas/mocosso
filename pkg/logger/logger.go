package logger

import (
	"io"
	"log/slog"
	"os"
	"sync"
)

var (
	logger     Logger
	loggerOnce sync.Once

	logWriter io.Writer
)

func SetLogWriter(out io.Writer) {
	logWriter = out
}

// getLogger returns a logger singleton configured to log in json format
func getLogger() Logger {
	loggerOnce.Do(func() {
		logger = NewSlogInterceptor(slog.New(slog.NewJSONHandler(logWriter, nil)))
	})

	return logger
}

// getLogWriter returns a writer for a logger or os.Stdout
func getLogWriter() io.Writer {
	if logWriter == nil {
		return os.Stdout
	}

	return logWriter
}

func Debug(msg string, args ...any) {
	getLogger().Debug(msg, args...)
}

func Info(msg string, args ...any) {
	getLogger().Info(msg, args...)
}

func Warn(msg string, args ...any) {
	getLogger().Warn(msg, args...)
}

func Error(msg string, args ...any) {
	getLogger().Error(msg, args...)
}
