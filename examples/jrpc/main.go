package main

import (
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	url := "https://ethereum-rpc.publicnode.com"
	blknum, err := jrpc.BlockNumber(url, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(blknum)
	fmt.Println(jrpc.BigIntToHexString(blknum))

	result, err := jrpc.GetBlockByNumber(url, 1, "latest", true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Number)
}
