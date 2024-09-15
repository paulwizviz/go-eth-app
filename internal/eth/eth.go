package eth

import (
	"log"
)

// Transaction is a structure for storing transaction data
type Transaction struct {
	Hash  string
	From  string
	To    string
	Value string
	Block int
}

// Block is a representation of a block from Ethereum
// node
type Block struct {
	Number       string        `json:"number"`
	Transactions []Transaction `json:"transactions"`
}

// BlockTxn is a replication of Block but
// replaced with the field BlockNum to make
// it easier to map for downstream operation.
type BlockTxn struct {
	BlockNum int
	Txns     []Transaction
}

// ReadNetwork is an operation to read data from
// the Ethereum network and ensure data is channelled
// to receiver.
func ReadNetwork(url string) chan BlockTxn {
	c := make(chan BlockTxn, 1)
	go func(ch chan BlockTxn) {
		for {
			bt := BlockTxn{}
			blockNumber, err := currentBlock(url)
			if err != nil {
				log.Println(err)
				continue
			}
			bt.BlockNum = int(blockNumber.Int64())
			txns, err := getBlockTransactions(url, blockNumber)
			if err != nil {
				log.Println(err)
				continue
			}
			bt.Txns = txns
			ch <- bt
		}
	}(c)
	return c
}

// LatestParseBlock represent a persistent store
// of the latest block ID.
type LatestParseBlock interface {
	Update(id int)
	GetID() int
}

// Parser represents a handler to enable a
// client application to obtained data from
// local data store.
type Parser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []Transaction
}
