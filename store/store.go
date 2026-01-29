package store

type Entry struct {
	Level   string
	Message string
	Time    string
}

type Store interface {
	Add(Entry)
	Get() []Entry
}
