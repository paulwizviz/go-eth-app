package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"paulwizviz/go-eth-app/internal/contract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deployContract(client *ethclient.Client, privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(3000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Smart contract bytecode and ABI (compiled using solc or Remix IDE)
	compiledContract := "0x..." // Bytecode
	tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, common.FromHex(compiledContract))

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contract deployed! Transaction Hash: %s\n", signedTx.Hash().Hex())

	// Get contract address
	contractAddress := crypto.CreateAddress(fromAddress, nonce)
	fmt.Println("Contract Address:", contractAddress.Hex())
	return contractAddress
}

func main() {

	// Private key from prefunded account
	privateKey, err := crypto.HexToECDSA("4c0883a69102937d6231471b5dbb6204fe512961708279a8911b138d9d808759")
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(3_000_000)

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
	addr, err := contract.DeployContract(context.Background(), client, privateKey, gasLimit, content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contract Address: ", addr)

}
