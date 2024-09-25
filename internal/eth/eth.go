package eth

import (
	"context"
	"log"
	"time"
)

// Transaction is a structure for storing transaction data
type Transaction struct {
	BlockHash            string `json:"blockHash"`
	Hash                 string `json:"hash"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Input                string `json:"input"`
	Value                string `json:"value"`
	Block                string `json:"blockNumber"`
	Type                 string `json:"type"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
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
	BlockNum string
	Txns     []Transaction
}

// ReadNetwork is an operation to read data from
// the Ethereum network and ensure data is channelled
// to receiver.
func ReadNetwork(c context.Context, url string) chan BlockTxn {
	ch := make(chan BlockTxn, 1)
	ticker := time.NewTicker(5000 * time.Millisecond)
	go func(ch chan BlockTxn) {
		for {
			select {
			case <-ticker.C:
				getLatestBlock(ch, url)
			case <-c.Done():
				return
			}
		}
	}(ch)

	getLatestBlock(ch, url)

	return ch
}

func getLatestBlock(ch chan BlockTxn, url string) {
	log.Println("Getting latest block...")
	bt := BlockTxn{}
	blockNumber, err := currentBlock(url)
	if err != nil {
		log.Println(err)
		return
	}
	bt.BlockNum = blockNumber.String()
	log.Printf("Got block %s", bt.BlockNum)
	txns, err := getBlockTransactions(url, blockNumber)
	if err != nil {
		log.Println(err)
		return
	}
	bt.Txns = txns
	ch <- bt
}
