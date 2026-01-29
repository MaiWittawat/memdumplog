package memdumplog

import (
	"errors"

	"github.com/MaiWittawat/memdumplog/adapter"
	"github.com/MaiWittawat/memdumplog/store"
)

type Logger struct {
	store store.Store
}

func New(cfg Config) (*Logger, error) {
	if cfg.BufferSize <= 0 {
		cfg.BufferSize = 100
	}

	st := store.NewMemory(cfg.BufferSize)

	switch cfg.Driver {
	case Logrus:
		adapter.UseLogrus(st)
	case Zap:
		adapter.UseZap(st)
	case Zerolog:
		adapter.UseZerolog(st)
	case Slog:
		adapter.UseSlog(st)
	default:
		return nil, errors.New("unsupported logger driver")
	}

	return &Logger{store: st}, nil
}

func (l *Logger) Logs() []store.Entry {
	return l.store.Get()
}
