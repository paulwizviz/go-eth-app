package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"paulwizviz/go-eth-app/internal/contract"
	"paulwizviz/go-eth-app/internal/jrpc"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	url := "http://localhost:8545"
	timeout := 50 * time.Millisecond

	// Get random dev account
	accts, err := jrpc.Accounts(context.TODO(), timeout, url, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accts)

	// Obtain the balance of random dev account
	bal, err := jrpc.GetBalance(context.TODO(), timeout, url, 1, accts[0], jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bal)

	// Private key from a hardcoded Hex.
	privateKey, err := crypto.HexToECDSA("4c0883a69102937d6231471b5dbb6204fe512961708279a8911b138d9d808759")
	if err != nil {
		log.Fatal(err)
	}

	// Get the address of the assigned to the private key
	pubkey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(pubkey).Hex()
	fmt.Println(address)

	// Transfer ether from dev account to the address associated with the hard coded
	// private key
	txn1 := jrpc.TxnArg{
		To:       address,
		From:     accts[0],
		Value:    "0x16345785d8a0000",
		Gas:      "0x5208",     // 21000
		GasPrice: "0x3b9aca00", // 1,000,000,000
	}

	txnHash, err := jrpc.SendTransaction(context.TODO(), timeout, url, 1, txn1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Txn Hash: ", txnHash)

	// Extract contract compuled data
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	content, err := contract.ExtractContent(fmt.Sprintf("%s/solidity/abi/hello/HelloWorld.bin", pwd))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)

	// Get the nounce from pending blocks
	nonce, err := jrpc.GetTxnCount(context.TODO(), timeout, url, 1, address, jrpc.BlockTagPENDING)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce", nonce)

	// Get suggested gas price
	gasPrice, err := jrpc.GasPrice(context.TODO(), timeout, url, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Gas price: ", gasPrice)

	// Get Chain ID
	chainID, err := jrpc.NetworkID(context.TODO(), timeout, url, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainid ", chainID)

	// Create a transaction embedding the contract content
	txn2 := contract.CreateContractEIP1559Txn(chainID.Int64(), uint64(nonce.Int64()), big.NewInt(1_000_000_000), gasPrice, 97590, []byte(content))
	signedTxn, err := contract.SignTransaction(txn2, uint64(chainID.Int64()), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	b, err := signedTxn.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	// Get signed transaction in hex string
	s := fmt.Sprintf("0x%v", hex.EncodeToString(b))

	// Send signed transaction to Dev node
	txnHash, err = jrpc.SendRawTransaction(context.TODO(), timeout, url, 1, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txnHash)
}
