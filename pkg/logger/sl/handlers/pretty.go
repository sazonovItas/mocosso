package slhandlers

import (
	"context"
	"encoding/json"
	"io"
	stdLog "log"
	"log/slog"

	"github.com/fatih/color"
)

// PrettyHandlerOptions struct of pretty handler options.
type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

// PrettyHandler struct of pretty handler.
type PrettyHandler struct {
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

// NewPrettyHandler method returns new pretty handler to log at given out writer.
func (opts PrettyHandlerOptions) NewPrettyHandler(out io.Writer) *PrettyHandler {
	return &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}
}

// Enabled method is implementation of slog handler interface.
func (h *PrettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

// Handle method is implementation of slog handler interface.
func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := "[" + r.Level.String() + "]"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var (
		b   []byte
		err error
	)
	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	time := color.MagentaString(r.Time.Format("[02/01 2006 15:04:05.000]"))
	msg := color.CyanString(r.Message)

	h.l.Println(time, level, msg, color.WhiteString(string(b)))

	return nil
}

// WithAttrs method is implementation of slog handler interface.
func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler.WithAttrs(attrs),
		l:       h.l,
		attrs:   attrs,
	}
}

// WithGroup method is implementation of slog handler interface.
func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
