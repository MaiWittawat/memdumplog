package adapter

import (
	"io"
	"time"

	"github.com/MaiWittawat/memdumplog/store"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type ZeroWriter struct {
	store store.Store
}

func NewZerolog(store store.Store) *ZeroWriter {
	return &ZeroWriter{store: store}
}

func UseZerolog(store store.Store) {
	mw := zerolog.MultiLevelWriter(
		zlog.Logger,
		&ZeroWriter{store: store},
	)

	zlog.Logger = zerolog.New(mw).With().Timestamp().Logger()
}

func (w *ZeroWriter) Write(p []byte) (n int, err error) {
	w.store.Add(store.Entry{
		Level:   "unknown",
		Message: string(p),
		Time:    time.Now().Format(time.RFC3339),
	})
	return len(p), nil
}

func NewZeroLogger(out io.Writer, store store.Store) zerolog.Logger {
	mw := zerolog.MultiLevelWriter(out, NewZerolog(store))
	return zerolog.New(mw).With().Timestamp().Logger()
}
