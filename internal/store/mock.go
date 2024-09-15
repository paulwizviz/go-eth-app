package store

import (
	"log"
	"sync"
)

type mockStorage struct {
	kvstore map[string][][]byte
	mu      sync.Mutex
}

func (s *mockStorage) Persists(key string, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, found := s.kvstore[key]
	if !found {
		v := [][]byte{}
		v = append(v, value)
		s.kvstore[key] = v
		log.Println("Append value", s.kvstore)
		return nil
	}
	v := s.kvstore[key]
	v = append(v, value)
	s.kvstore[key] = v

	log.Println("New value", s.kvstore)

	return nil

}

func (s *mockStorage) Get(key string) ([][]byte, error) {
	v, found := s.kvstore[key]
	if !found {
		return nil, ErrValueNotFound
	}
	return v, nil
}
