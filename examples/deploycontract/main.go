package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"paulwizviz/go-eth-app/internal/contract"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	// Private key from prefunded account
	privateKey, err := crypto.HexToECDSA("4c0883a69102937d6231471b5dbb6204fe512961708279a8911b138d9d808759")
	if err != nil {
		log.Fatal(err)
	}

	pubkey := privateKey.PublicKey

	address := crypto.PubkeyToAddress(pubkey).Hex()
	fmt.Println(address)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	content, err := contract.ExtractContent(fmt.Sprintf("%s/solidity/abi/hello/HelloWorld.bin", pwd))
	if err != nil {
		log.Fatal(err)
	}
	// Connect to Geth local node
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Geth local node")

	// Deploy contract
	gasLimit := uint64(138_612)
	addr, err := contract.DeployContract(context.Background(), client, privateKey, 1_000_000_000, 1_000_000_000, gasLimit, content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contract Address: ", addr)

}
