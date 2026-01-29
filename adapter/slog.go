package adapter

import (
	"context"
	"github.com/MaiWittawat/memdumplog/store"
	"time"

	"log/slog"
)

type SlogHandler struct {
	store store.Store
}

func NewSlogHandler(store store.Store) slog.Handler {
	return &SlogHandler{store: store}
}

func UseSlog(store store.Store) {
	slog.SetDefault(
		slog.New(NewSlogHandler(store)),
	)
}

func (h *SlogHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *SlogHandler) Handle(_ context.Context, r slog.Record) error {
	h.store.Add(store.Entry{
		Level:   r.Level.String(),
		Message: r.Message,
		Time:    time.Now().Format(time.RFC3339),
	})
	return nil
}

func (h *SlogHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *SlogHandler) WithGroup(_ string) slog.Handler {
	return h
}
