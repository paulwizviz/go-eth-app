package main

import (
	"context"
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	url := "https://ethereum-rpc.publicnode.com"
	blknum, err := jrpc.BlockNumber(context.TODO(), url, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(blknum)
	fmt.Println(jrpc.BigIntToHexString(blknum))

	block, err := jrpc.GetBlockByNumber(context.TODO(), url, 1, "latest", true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Block number: %v\n", block.Number)
	fmt.Printf("Block transactions. Types: %T", block.Transactions)
}
