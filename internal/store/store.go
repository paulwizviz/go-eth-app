package store

import (
	"errors"
)

var (
	ErrEntryExists   = errors.New("entry exists error")
	ErrValueNotFound = errors.New("value not found")
)

// Storage is an abstraction of a persistent key value store
type Storage interface {
	Persists(key string, value []byte) error
	Get(key string) ([][]byte, error)
}

// NewMockStorage instantage a mock storage
func NewMockStorage() Storage {
	return &mockStorage{
		kvstore: make(map[string][][]byte),
	}
}
