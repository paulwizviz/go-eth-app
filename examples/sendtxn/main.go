package main

import (
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	url := "http://localhost:8545"
	addr := "0x1a12FAD7cA2fb46417Ffc3dE3Cda72219459292C" // Dev mode address

	t := jrpc.TxnArg{
		To:       addr,
		From:     "0x8793eF7f91dB0891E22DB3e1F8C0Fe616673a38C",
		Value:    "0x16345785d8a0000",
		Gas:      "0x5208",
		GasPrice: "0x3b9aca00",
	}

	result, err := jrpc.SendTransaction(url, 1, t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	bal, err := jrpc.GetBalance(url, 1, addr, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bal)
}
