package main

import (
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/jrpc"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	url := "http://localhost:8545"
	accts, err := jrpc.Accounts(url, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accts)

	bal, err := jrpc.GetBalance(url, 1, accts[0], jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bal)

	// Private key from prefunded account
	privateKey, err := crypto.HexToECDSA("4c0883a69102937d6231471b5dbb6204fe512961708279a8911b138d9d808759")
	if err != nil {
		log.Fatal(err)
	}

	pubkey := privateKey.PublicKey

	address := crypto.PubkeyToAddress(pubkey).Hex()
	fmt.Println(address)

	txn := jrpc.TxnArg{
		To:       address,
		From:     accts[0],
		Value:    "0x16345785d8a0000",
		Gas:      "0x5208",
		GasPrice: "0x3b9aca00",
	}

	txnHash, err := jrpc.SendTransaction(url, 1, txn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Txn Hash: ", txnHash)

	accts, err = jrpc.Accounts(url, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accts)

	bal, err = jrpc.GetBalance(url, 1, address, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bal)
}
