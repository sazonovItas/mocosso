package slhandlers

import (
	"context"
	"encoding/json"
	"io"
	stdLog "log"
	"log/slog"
)

const (
	timeKey    string = "timestamp"
	levelKey   string = "level"
	messageKey string = "msg"
)

// DiscardHandlerOptions struct of discard handler options.
type DiscardHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

// DiscardHandler struct of discard handler.
type DiscardHandler struct {
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

// NewDiscardHandler method returns new discard handler to log at given out writer.
func (opts *DiscardHandlerOptions) NewDiscardHandler(out io.Writer) *DiscardHandler {
	return &DiscardHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}
}

// Enabled method returns enabled level or not to log.
func (h *DiscardHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

// Handle method is implementation of slog handler interface.
func (h *DiscardHandler) Handle(ctx context.Context, r slog.Record) error {
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	fields[levelKey] = r.Level
	fields[timeKey] = r.Time
	fields[messageKey] = r.Message

	var (
		b   []byte
		err error
	)
	if len(fields) > 0 {
		b, err = json.Marshal(fields)
		if err != nil {
			return err
		}
	}

	h.l.Println(string(b))

	return nil
}

// WithAttrs method is implementation of slog handler interface.
func (h *DiscardHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &DiscardHandler{
		Handler: h.Handler.WithAttrs(attrs),
		l:       h.l,
		attrs:   attrs,
	}
}

// WithGroup method is implementation of slog handler interface.
func (h *DiscardHandler) WithGroup(name string) slog.Handler {
	return &DiscardHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
