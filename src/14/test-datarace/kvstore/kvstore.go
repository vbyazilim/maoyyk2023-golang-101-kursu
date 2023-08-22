package kvstore

import "errors"

var errKeyNotFound = errors.New("key not found")

// Store is key-value store!
type Store struct {
	db map[string]string
}

// Set new key to store.
func (s *Store) Set(k, v string) error {
	s.db[k] = v
	return nil
}

// Get accepts key, returns value and error.
func (s *Store) Get(k string) (string, error) {
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
