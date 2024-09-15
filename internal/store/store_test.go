package store

import "fmt"

func Example_mockStorage() {
	s := NewMockStorage()

	s.Persists("abc", []byte("Hello"))
	s.Persists("abc", []byte("efc"))

	v, _ := s.Get("abc")
	fmt.Println(v)

	// Output:
	// [[72 101 108 108 111] [101 102 99]]
}
