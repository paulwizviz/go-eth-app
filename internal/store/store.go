package store

import (
	"errors"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

// Storage is an abstraction of a persistent key value store
type Storage interface {
	Append(key string, value []byte) error
	Get(key string) ([][]byte, error)
	Set(key string, value [][]byte) error
	Keys() []string
}
