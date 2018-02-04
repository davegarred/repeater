package persist

import (
	"sync"
)

type MemStore struct {
	items map[string]string
	mimetype map[string]string
	mu    sync.Mutex
}

func NewMemStore() *MemStore {
	return &MemStore{items: make(map[string]string), mimetype: make(map[string]string)}
}

func (s *MemStore) Store(mimetype string, k string, v string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if curVal := s.items[k]; curVal != "" {
		return KEY_CONFLICT
	}
	s.items[k] = v
	s.mimetype[k] = mimetype
	return nil
}
func (s *MemStore) Retrieve(key string) (string, error) {
	return s.items[key], nil
}

func (s *MemStore) Delete(k string) error {
	delete(s.items, k)
	return nil
}
