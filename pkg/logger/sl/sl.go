package slogger

import (
	"fmt"
	"io"
	"log/slog"

	slhandlers "github.com/sazonovItas/auth-service/pkg/logger/sl/handlers"
)

const (
	ConsoleEncoding string = "console"
	JSONEncoding    string = "json"
)

// NewLogger functions creates new logger depence on logger type json and text.
func MustNewLogger(encoding string, level slog.Level, out io.Writer) (log *slog.Logger) {
	switch encoding {
	case JSONEncoding:
		log = NewDiscardLogger(level, out)
	case ConsoleEncoding:
		log = NewConsoleLogger(level, out)
	default:
		panic(fmt.Errorf("unknow type of encoding"))
	}

	return log
}

// NewPrettyLogger function creates json pretty handler for slog logger.
func NewPrettyLogger(level slog.Level, out io.Writer) *slog.Logger {
	opts := slhandlers.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewPrettyHandler(out)
	return slog.New(handler)
}

// NewDiscardLogger function creates json discard handler for slog logger.
func NewDiscardLogger(level slog.Level, out io.Writer) *slog.Logger {
	opts := slhandlers.DiscardHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewDiscardHandler(out)
	return slog.New(handler)
}

// NewTextLogger function creates text handler for slog logger.
func NewConsoleLogger(level slog.Level, out io.Writer) *slog.Logger {
	opts := slhandlers.TextHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewTextHandler(out)
	return slog.New(handler)
}

// Err function returns slog attribute from error.
func Err(err error) slog.Attr {
	if err == nil {
		slog.String("error", "nil")
	}

	return slog.String("error", err.Error())
}
