package main

import (
	"context"
	"fmt"
	"log"

	"github.com/paulwizviz/go-eth-app/internal/jrpc"
)

func main() {

	client := jrpc.NewDefaultClient("http://localhost:8545")
	receipt, err := client.GetTxnReceipt(context.TODO(), 1, "0x75367bbc0944507f8ea0ef17d40086e008ea9c863b4d3854fc2a83c8e85c06ae")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.ContractAddress)
}
