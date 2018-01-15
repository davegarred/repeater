package persist

import (
	"sync"

	"github.com/google/uuid"
)

type Store interface {
	Store(string) uuid.UUID
	Retrieve(uuid.UUID) string
}

type MemStore struct {
	items map[uuid.UUID]string
	mu    sync.Mutex
}

func NewStore() *MemStore {
	return &MemStore{items: make(map[uuid.UUID]string)}
}

func (s *MemStore) Store(v string) uuid.UUID {
	s.mu.Lock()
	key := uuid.New()
	s.items[key] = v
	s.mu.Unlock()
	return key
}
func (s *MemStore) Retrieve(key uuid.UUID) string {
	return s.items[key]
}
