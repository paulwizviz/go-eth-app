package main

import (
	"context"
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/jrpc"
)

func main() {

	client := jrpc.NewDefaultClient("http://localhost:8545")

	number, err := client.BlockNumber(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(number)

	blk, err := client.GetBlockByNumber(context.TODO(), 1, fmt.Sprintf("0x%x", number), true)
	if err != nil {
		log.Fatal(err)
	}

	var txnHash string
	switch v := blk.Transactions.(type) {
	case []jrpc.TxnEIP2718:
		for _, tx := range v {
			txnHash = tx.Hash
			fmt.Println(txnHash)
		}
	}

	receipt, err := client.GetTxnReceipt(context.TODO(), 1, txnHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-->", receipt)
}
