package persist

import (
	"sync"

	"github.com/google/uuid"
)

type Store struct {
	items map[uuid.UUID]string
	mu    sync.Mutex
}

func NewStore() *Store {
	return &Store{items: make(map[uuid.UUID]string)}
}

func (s *Store) Store(v string) uuid.UUID {
	s.mu.Lock()
	key := uuid.New()
	s.items[key] = v
	s.mu.Unlock()
	return key
}
func (s *Store) Retrieve(key uuid.UUID) string {
	return s.items[key]
}
