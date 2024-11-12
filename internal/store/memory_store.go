package store

import (
	"sort"
	"strings"
	"sync"
)

// NewInMemoryStorage instantiates a new in-memory storage
func NewInMemoryStorage() Storage {
	return &InMemoryStorage{
		data: make(map[string][][]byte),
	}
}

// InMemoryStorage is an in-memory store of key/list of values
type InMemoryStorage struct {
	data map[string][][]byte
	mu   sync.RWMutex
}

// Append appends a new byte slice to "key"
func (s *InMemoryStorage) Append(key string, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, found := s.data[key]
	if !found {
		s.data[key] = [][]byte{}
	}
	v := s.data[key]
	v = append(v, value)
	s.data[key] = v

	return nil
}

// Get gets the slice of byte slices for a given key
func (s *InMemoryStorage) Get(key string) ([][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, found := s.data[key]
	if !found {
		return nil, ErrKeyNotFound
	}
	return v, nil
}

// Set sets the value of "key"
func (s *InMemoryStorage) Set(key string, value [][]byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return nil
}

// Keys returns a list of all keys in the store
func (s *InMemoryStorage) Keys() []string {
	keys := []string{}

	s.mu.RLock()
	defer s.mu.RUnlock()

	for k := range s.data {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) < 0
	})

	return keys
}
