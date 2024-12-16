package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"paulwizviz/go-eth-app/internal/contract"
	"paulwizviz/go-eth-app/internal/jrpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	client := jrpc.NewDefaultClient("http://localhost:8545")

	// Get random dev account
	devAcct, err := getDevRandomAcct(client)
	if err != nil {
		log.Fatal(err)
	}

	targetAddr, privKey, err := getTargetAddr()
	if err != nil {
		log.Fatal(err)
	}

	txnHash, err := transferToTargetAddr(client, targetAddr, devAcct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Txn Hash: ", txnHash)

	// Get the nounce from pending blocks
	nonce, err := client.GetTxnCount(context.TODO(), 1, targetAddr, jrpc.BlockTagPENDING)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce", nonce)

	// Get suggested gas price
	gasPrice, err := client.GasPrice(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Gas price: ", gasPrice)

	// Get Chain ID
	chainID, err := client.NetworkID(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainid ", chainID)

	// Extract contract compuled data
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	binContent, err := contract.ExtractContentBin(fmt.Sprintf("%s/solidity/abi/hello/HelloWorld.bin", pwd))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(binContent)

	abiContent, err := contract.ExtractContractABI(fmt.Sprintf("%s/solidity/abi/hello/HelloWorld.abi", pwd))
	if err != nil {
		log.Fatal(err)
	}

	data, err := contractData(binContent, abiContent)
	if err != nil {
		log.Fatal(err)
	}

	// Create a transaction embedding the contract content
	txn2 := contract.CreateContractEIP1559Txn(chainID.Int64(), uint64(nonce.Int64()), big.NewInt(1_000_000_000), gasPrice, 97590, data)
	signedTxn, err := contract.SignTransaction(txn2, uint64(chainID.Int64()), privKey)
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
	txnHash2, err := client.SendRawTransaction(context.TODO(), 1, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("->", txnHash2)
}

func getDevRandomAcct(client jrpc.Client) (string, error) {
	accts, err := client.Accounts(context.TODO(), 1)
	if err != nil {
		return "", err
	}
	devAcct := accts[0]

	bal, err := client.GetBalance(context.TODO(), 1, devAcct, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bal)

	return devAcct, nil
}

func getTargetAddr() (string, *ecdsa.PrivateKey, error) {
	// Private key from a hardcoded Hex.
	privateKey, err := crypto.HexToECDSA("4c0883a69102937d6231471b5dbb6204fe512961708279a8911b138d9d808759")
	if err != nil {
		return "", nil, err
	}

	// Get the address of the assigned to the private key
	pubkey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(pubkey).Hex()
	return address, privateKey, nil
}

func transferToTargetAddr(client jrpc.Client, targetAddr string, devAcct string) (string, error) {
	// Transfer ether from dev account to the address associated with the hard coded
	// private key
	txn1 := jrpc.TxnArg{
		To:       targetAddr,
		From:     devAcct,
		Value:    "0x16345785d8a0000", // 100,000,000,000,000,000
		Gas:      "0x5208",            // 21000
		GasPrice: "0x3b9aca00",        // 1,000,000,000
	}

	txnHash, err := client.SendTransaction(context.TODO(), 1, txn1)
	if err != nil {
		return "", err
	}
	return txnHash, nil
}

func contractData(contractBin string, contractABI string) ([]byte, error) {
	initialValue := big.NewInt(1_000)
	constructorArg, err := contract.EncodeConstructorArg(contractABI, initialValue)
	if err != nil {
		return nil, err
	}
	data := append(common.FromHex(contractBin), constructorArg...)
	return data, nil
}
