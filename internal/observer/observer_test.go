package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	o := New()

	results1 := []string{}
	results2 := []string{}
	results3 := []string{}

	fmt.Printf("starting len: %d\n", len(results1))

	s1 := o.Subscribe("topic-a")
	s2 := o.Subscribe("topic-a")
	s3 := o.Subscribe("topic-b")

	end := make(chan bool, 1)

	go func(o *Observer) {
		o.Notify("topic-a", []byte("a"))
		o.Notify("topic-b", []byte("b"))
		o.Notify("topic-a", []byte("c"))
		s1.Unsubscribe()
		o.Notify("topic-a", []byte("d"))

		time.Sleep(100 * time.Millisecond)
		end <- true
	}(o)

outer:
	for {
		select {
		case m, ok := <-s1.Ch:
			if !ok {
				break
			}
			results1 = append(results1, string(m))
		case m, ok := <-s2.Ch:
			if !ok {
				break
			}
			results2 = append(results2, string(m))
		case m, ok := <-s3.Ch:
			if !ok {
				break
			}
			results3 = append(results3, string(m))
		case <-end:
			break outer
		}
	}

	if len(results1) != 2 {
		t.Errorf("results1: expected length of 2; got %d", len(results1))
	}
	if len(results2) != 3 {
		t.Errorf("results2: expected length of 3; got %d", len(results2))
	}
	if len(results3) != 1 {
		t.Errorf("results3: expected length of 1; got %d", len(results3))
	}
}
