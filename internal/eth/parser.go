package eth

import (
	"encoding/json"
	"log"
	"txparser/internal/counter"
	"txparser/internal/observer"
	"txparser/internal/store"
)

// Parser represents a handler to enable a
// client application to obtained data from
// local data store.
type Parser interface {
	// last parsed block
	GetCurrentBlock() string
	// add address to observer
	Subscribe(address string) *observer.Subscription
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []Transaction
	// GetAddresses returns a list of all addresses seen
	GetAddresses() []string
	// GetCount returns the tx count for a given address
	GetCount(address string) int64
}

func NewDefaultParser(blocktxn chan BlockTxn) Parser {
	d := defaultParser{
		latestBlock: NewLatestParseBlock(),
		txnStorage:  store.NewInMemoryStorage(),
		observer:    observer.New(),
		counter:     counter.New(),
	}
	// Initiate a Goroutine to read data
	// from the Ethereum network.
	go func() {
		for b := range blocktxn {
			if d.latestBlock.Get() >= b.BlockNum {
				log.Printf("Already processed block %s", b.BlockNum)
				continue
			}
			d.latestBlock.Update(b.BlockNum)

			for _, tx := range b.Txns {
				txMarshal, err := json.Marshal(tx)
				if err != nil {
					log.Println(err)
					continue
				}

				d.counter.Add(tx.From)
				d.counter.Add(tx.To)
				d.observer.Notify(tx.From, txMarshal)
				d.observer.Notify(tx.To, txMarshal)
				if err := d.txnStorage.Append(tx.From, txMarshal); err != nil {
					log.Println(err)
				}
				if err := d.txnStorage.Append(tx.To, txMarshal); err != nil {
					log.Println(err)
				}
			}
		}
	}()
	return &d
}

type defaultParser struct {
	latestBlock LatestParseBlock   // persistent store for latest block
	txnStorage  store.Storage      // store for transactions
	observer    *observer.Observer // subscriber list
	counter     *counter.Counter
}

func (d *defaultParser) GetCurrentBlock() string {
	return d.latestBlock.Get()
}

func (d *defaultParser) Subscribe(address string) *observer.Subscription {
	return d.observer.Subscribe(address)
}

func (d *defaultParser) GetTransactions(address string) []Transaction {
	txs, err := d.txnStorage.Get(address)
	if err != nil {
		log.Println(err)
		return nil
	}

	var txns []Transaction
	for _, tx := range txs {
		var t Transaction
		if err := json.Unmarshal(tx, &t); err != nil {
			continue
		}
		txns = append(txns, t)
	}
	return txns
}

func (d *defaultParser) GetAddresses() []string {
	return d.txnStorage.Keys()
}

func (d *defaultParser) GetCount(address string) int64 {
	return d.counter.Get(address)
}
