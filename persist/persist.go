package persist

import (
	"sync"

	"github.com/google/uuid"
)

type Store interface {
	Store(string) string
	Retrieve(string) string
}

type MemStore struct {
	items map[string]string
	mu    sync.Mutex
}

func NewStore() *MemStore {
	return &MemStore{items: make(map[string]string)}
}

func (s *MemStore) Store(v string) string {
	s.mu.Lock()
	key := uuid.New().String()
	s.items[key] = v
	s.mu.Unlock()
	return key
}
func (s *MemStore) Retrieve(key string) string {
	return s.items[key]
}
