package store

import "sync"

type MemoryStore struct {
	mu    sync.Mutex
	buf   []Entry
	limit int
}

func NewMemory(limit int) *MemoryStore {
	return &MemoryStore{
		buf:   make([]Entry, 0, limit),
		limit: limit,
	}
}

func (m *MemoryStore) Add(e Entry) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.buf = append(m.buf, e)
	if len(m.buf) > m.limit {
		m.buf = m.buf[1:]
	}
}

func (m *MemoryStore) Get() []Entry {
	m.mu.Lock()
	defer m.mu.Unlock()

	out := make([]Entry, len(m.buf))
	copy(out, m.buf)
	return out
}
