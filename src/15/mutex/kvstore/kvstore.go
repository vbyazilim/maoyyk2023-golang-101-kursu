package kvstore

import (
	"errors"
	"sync"
)

var errKeyNotFound = errors.New("key not found")

// Store is key-value store!
type Store struct {
	mu sync.RWMutex
	db map[string]string
}

// Set new key to store.
func (s *Store) Set(k, v string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.db[k] = v
	return nil
}

// Get accepts key, returns value and error.
func (s *Store) Get(k string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.db[k]
	if !ok {
		return "", errKeyNotFound
	}
	return v, nil
}

// New returns new Store instance.
func New(db map[string]string) Store {
	return Store{db: db}
}
