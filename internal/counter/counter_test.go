package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	wg := sync.WaitGroup{}
	counter := New()

	for range 100 {
		wg.Add(1)
		counter.Add("test")
		wg.Done()
	}

	wg.Wait()

	count := counter.Get("test")
	if count != 100 {
		t.Errorf("expected 100; got %d", count)
	}
}
