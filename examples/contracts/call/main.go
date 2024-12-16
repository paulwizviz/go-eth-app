package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"paulwizviz/go-eth-app/internal/contract"
	"paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	client := jrpc.NewDefaultClient("http://localhost:8545")

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	abiContent, err := contract.ExtractContractABI(fmt.Sprintf("%s/solidity/abi/hello/HelloWorld.abi", pwd))
	if err != nil {
		log.Fatal(err)
	}

	data, err := contract.EncodeFuncCall(abiContent, "getValue")
	if err != nil {
		log.Fatal(err)
	}

	txnArg := jrpc.TxnArg{
		To:   "0xC2999FE1f8c6E506bEc4c687f7c638069BbC16bC",
		Data: fmt.Sprintf("0x%v", hex.EncodeToString(data)),
	}

	txnHash, err := client.Call(context.TODO(), 1, txnArg, jrpc.BlockTagFinalized)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txnHash)

}
