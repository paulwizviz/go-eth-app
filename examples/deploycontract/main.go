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

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	// Get random dev account
	url := "http://localhost:8545"

	accts, err := jrpc.Accounts(context.TODO(), url, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accts)

	bal, err := jrpc.GetBalance(context.TODO(), url, 1, accts[0], jrpc.BlockTagLATEST)
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

	// Transfer ether from dev to address
	txn1 := jrpc.TxnArg{
		To:       address,
		From:     accts[0],
		Value:    "0x16345785d8a0000",
		Gas:      "0x5208",     // 21000
		GasPrice: "0x3b9aca00", // 1,000,000,000
	}

	txnHash, err := jrpc.SendTransaction(context.TODO(), url, 1, txn1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Txn Hash: ", txnHash)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	content, err := contract.ExtractContent(fmt.Sprintf("%s/solidity/abi/hello/HelloWorld.bin", pwd))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)

	nonce, err := jrpc.GetTxnCount(context.TODO(), url, 1, address, jrpc.BlockTagPENDING)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce", nonce)

	gasPrice, err := jrpc.GasPrice(context.TODO(), url, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Gas price: ", gasPrice)

	chainID, err := jrpc.NetworkID(context.TODO(), url, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainid ", chainID)

	txn2 := contract.CreateContractEIP1559Txn(chainID.Int64(), uint64(nonce.Int64()), big.NewInt(1_000_000_000), gasPrice, 97590, []byte(content))
	signedTxn, err := contract.SignTransaction(txn2, uint64(chainID.Int64()), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	b, err := signedTxn.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	s := fmt.Sprintf("0x%v", hex.EncodeToString(b))

	txnHash, err = jrpc.SendRawTransaction(context.TODO(), url, 1, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txnHash)
}
