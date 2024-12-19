package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	// Instantiate a default client of internal JSON-RPC package
	client := jrpc.NewDefaultClient("http://localhost:8545")

	// Get the most recent block
	number, err := client.BlockNumber(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}
	blknum := hexutil.EncodeBig(number)
	fmt.Println(number, blknum)

}
