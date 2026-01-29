package adapter

import (
	"github.com/MaiWittawat/memdumplog/store"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusHook struct {
	store store.Store
}

func NewLogrusHook(store store.Store) *LogrusHook {
	return &LogrusHook{store: store}
}

func UseLogrus(st store.Store) {
	logrus.AddHook(NewLogrusHook(st))
}

func (h *LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *LogrusHook) Fire(entry *logrus.Entry) error {
	h.store.Add(store.Entry{
		Level:   entry.Level.String(),
		Message: entry.Message,
		Time:    entry.Time.Format(time.RFC3339),
	})
	return nil
}
