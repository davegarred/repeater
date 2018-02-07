package persist

import (
	"sync"
)

// MemStore will store objects in memory only
type MemStore struct {
	items map[string]string
	mimetype map[string]string
	mu    sync.Mutex
}

// NewMemStore returns a new *MemStore
func NewMemStore() *MemStore {
	return &MemStore{items: make(map[string]string), mimetype: make(map[string]string)}
}

// Store will persist the object value and mimetype indexed by the key
func (s *MemStore) Store(mimetype string, k string, v string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if curVal := s.items[k]; curVal != "" {
		return KeyConflict
	}
	s.items[k] = v
	s.mimetype[k] = mimetype
	return nil
}


// Retrieve takes a key and returns the stored value or an error
func (s *MemStore) Retrieve(key string) (*StoredObject, error) {
	object := s.items[key]
	if len(object) == 0 {
		return nil, nil
	}
	storedObject := &StoredObject{"application/json", object}
	return storedObject, nil
}


// Delete removes the key-value pair associated with the given key
func (s *MemStore) Delete(key string) error {
	delete(s.items, key)
	return nil
}
