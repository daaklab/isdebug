package sloghandler

import (
	"context"
	"log/slog"

	"github.com/daaklab/isdebug"
)

type handler struct {
	delegate slog.Handler
}

var _ slog.Handler = new(handler)

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return isdebug.Enabled || h.delegate.Enabled(ctx, level)
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	return h.delegate.Handle(ctx, r)
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	delegate := h.delegate.WithAttrs(attrs)
	return New(delegate)
}

func (h *handler) WithGroup(name string) slog.Handler {
	delegate := h.delegate.WithGroup(name)
	return New(delegate)
}

func New(delegate slog.Handler) slog.Handler {
	return &handler{delegate: delegate}
}
