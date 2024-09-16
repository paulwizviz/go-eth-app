package eth

import (
	"sync"
)

// NewLatestParseBlock instantiate a parsed block counter
func NewLatestParseBlock() LatestParseBlock {
	return &latestBlock{
		id: -1,
	}
}

type latestBlock struct {
	mu sync.Mutex
	id int
}

func (l *latestBlock) Update(id int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.id = id
}

func (l *latestBlock) GetID() int {
	return l.id
}
