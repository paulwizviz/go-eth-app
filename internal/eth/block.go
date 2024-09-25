package eth

import (
	"sync"
)

// LatestParseBlock represent a persistent store
// of the latest block ID.
type LatestParseBlock interface {
	Update(block string)
	Get() string
}

// NewLatestParseBlock instantiate a parsed block counter
func NewLatestParseBlock() LatestParseBlock {
	return &latestBlock{
		block: "-1",
	}
}

type latestBlock struct {
	mu    sync.Mutex
	block string
}

func (l *latestBlock) Update(block string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.block = block
}

func (l *latestBlock) Get() string {
	return l.block
}
