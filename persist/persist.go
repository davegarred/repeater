package persist

import (
	"sync"
)

const KEY_CONFLICT = Error("Key conflict on store")

type Error string

func (e Error) Error() string {
	return string(e)
}

type Store interface {
	Store(string, string) error
	Retrieve(string) (string, error)
}

type MemStore struct {
	items map[string]string
	mu    sync.Mutex
}

func NewMemStore() *MemStore {
	return &MemStore{items: make(map[string]string)}
}

func (s *MemStore) Store(k string, v string) error {
	s.mu.Lock()
	if curVal := s.items[k]; curVal != "" {
		return KEY_CONFLICT
	}
	s.items[k] = v
	s.mu.Unlock()
	return nil
}
func (s *MemStore) Retrieve(key string) (string, error) {
	return s.items[key], nil
}
