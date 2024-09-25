package store

import (
	"bytes"
	"errors"
	"testing"
)

func TestInMemoryStorage(t *testing.T) {
	s := NewInMemoryStorage()

	// test Append
	s.Append("abc", []byte("Hello"))
	s.Append("abc", []byte("eth"))

	v, _ := s.Get("abc")

	if len(v) != 2 {
		t.Errorf("expected length of 2; got %d\n", len(v))
	}
	if !bytes.Equal(v[0], []byte("Hello")) {
		t.Errorf("expected v[0] to be []byte(Hello); got []byte(%s)", string(v[0]))
	}
	if !bytes.Equal(v[1], []byte("eth")) {
		t.Errorf("expected v[1] to be []byte(eth); got []byte(%s)", string(v[1]))
	}

	// Test ErrValueNotFound
	_, err := s.Get("foo")
	if !errors.Is(err, ErrKeyNotFound) {
		t.Errorf("expected \"key not found\"; got \"%s\"", err)
	}

	// test Set
	s.Set("abc", [][]byte{[]byte("Bye!")})

	v, _ = s.Get("abc")

	if len(v) != 1 {
		t.Errorf("expected length of 1; got %d\n", len(v))
	}
	if !bytes.Equal(v[0], []byte("Bye!")) {
		t.Errorf("expected v[0] to be []byte(Bye!); got []byte(%s)", string(v[0]))
	}

	// test Keys
	s.Set("bar", [][]byte{[]byte("test")})

	keys := s.Keys()

	if len(keys) != 2 {
		t.Errorf("expected length of 2; got %d", len(keys))
	}
	if keys[0] != "abc" {
		t.Errorf("expected keys[0] to be \"abc\"; got \"%s\"", keys[0])
	}
	if keys[1] != "bar" {
		t.Errorf("expected keys[1] to be \"bar\"; got \"%s\"", keys[1])
	}
}
