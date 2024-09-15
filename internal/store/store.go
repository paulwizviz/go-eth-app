package store

import (
	"errors"
)

var (
	ErrEntryExists   = errors.New("entry exists error")
	ErrValueNotFound = errors.New("value not found")
)

type Storage interface {
	Persists(key string, value []byte) error
	Get(key string) ([][]byte, error)
}

func NewMockStorage() Storage {
	return &mockStorage{
		kvstore: make(map[string][][]byte),
	}
}
