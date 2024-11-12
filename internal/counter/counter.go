package counter

import "sync"

// Counter tracks the number of seen txs for a given address
type Counter struct {
	sync.RWMutex
	counts map[string]int64
}

// NewCounter creates a new Counter instance
func New() *Counter {
	return &Counter{
		counts: map[string]int64{},
	}
}

// Add increments the count for a given topic
func (t *Counter) Add(topic string) {
	t.Lock()
	if _, found := t.counts[topic]; !found {
		t.counts[topic] = 0
	}
	t.counts[topic] += 1
	t.Unlock()
}

// Get gets the count for a given topic
func (t *Counter) Get(topic string) int64 {
	t.RLock()
	defer t.RUnlock()
	if _, found := t.counts[topic]; !found {
		return 0
	}
	return t.counts[topic]
}
