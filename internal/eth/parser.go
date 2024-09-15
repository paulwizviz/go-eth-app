package eth

import (
	"encoding/json"
	"log"
	"txparser/internal/store"
)

func NewDefaultParser(btc chan BlockTxn) Parser {
	d := defaultParser{
		latestBlock: NewLatestParseBlock(),
		txnStorage:  store.NewMockStorage(),
		subcribers:  make(map[string]bool),
	}
	// Initiate a Goroutine to read data
	// from the Ethereum network.
	go func() {
		for b := range btc {
			d.latestBlock.Update(b.BlockNum)
			for _, tx := range b.Txns {
				addr := tx.From
				d.subcribers[addr] = true
				txMarshal, err := json.Marshal(tx)
				if err != nil {
					log.Println(err)
					continue
				}
				if err := d.txnStorage.Persists(addr, txMarshal); err != nil {
					log.Println(err)
				}
			}
		}
	}()
	return &d
}

type defaultParser struct {
	latestBlock LatestParseBlock // persistent store for latest block
	txnStorage  store.Storage    // store for transaction
	subcribers  map[string]bool  // subscriber list
}

func (d *defaultParser) GetCurrentBlock() int {
	return d.latestBlock.GetID()
}

func (d *defaultParser) Subscribe(address string) bool {
	v, found := d.subcribers[address]
	if !found {
		return found
	}
	return v
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
