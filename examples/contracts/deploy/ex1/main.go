// Copyright 2024 The Contributors to go-eth-app
// This file is part of the go-eth-app project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the License.
//
// For a list of contributors, refer to the CONTRIBUTORS file or the
// repository's commit history.

// This example demonstrate the steps involved in deploying
// contracts to Geth node in dev mode using internal JSON-RPC.

package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/paulwizviz/go-eth-app/internal/contract"
	"github.com/paulwizviz/go-eth-app/internal/jrpc"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	// Establish a jprc clent.
	client := jrpc.NewDefaultClient("http://localhost:8545")

	devAcc, err := getDevRandomAcct(context.TODO(), client)
	if err != nil {
		log.Fatal(err)
	}

	// Extract the content of compiled solidity
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	contractBin, err := contract.ExtractContentBin(fmt.Sprintf("%s/solidity/abi/hello/HelloWorld.bin", pwd))
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := contract.ExtractContractABI(fmt.Sprintf("%s/solidity/abi/hello/HelloWorld.abi", pwd))
	if err != nil {
		log.Fatal(err)
	}

	// Convert ABI JSON into abi.ABI type
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}

	// Encode the constructor argument
	constructorArg := big.NewInt(100)
	encodedConstArg, err := parsedABI.Pack("", constructorArg)
	if err != nil {
		log.Fatal(err)
	}
	// Pack contract byte slice and encoded argument byte slice into a single byte slice
	fullData := append(common.FromHex(contractBin), encodedConstArg...)

	// Deploy the contract to the node
	txn := jrpc.TxnArg{
		From: devAcc,
		Data: fmt.Sprintf("0x%X", fullData),
		Gas:  "0x800000",
	}
	txnHash, err := client.SendTransaction(context.TODO(), 1, txn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txnHash)

	// Sleep to allow node to write to chain.
	time.Sleep(1 * time.Second)

	// Get receipt for the transaction hash
	receipt, err := client.GetTxnReceipt(context.TODO(), 1, txnHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.ContractAddress)

	// Encode the function "getValue" and invoke it
	encodedGetValue, err := parsedABI.Pack("getValue")
	if err != nil {
		log.Fatal(err)
	}

	getValTxn := jrpc.TxnArg{
		To:   receipt.ContractAddress,
		Data: fmt.Sprintf("0x%X", encodedGetValue),
	}
	getValueCall1, err := client.Call(context.TODO(), 1, getValTxn, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getValueCall1)
	value1 := new(big.Int)
	value1.SetString(getValueCall1[2:], 16)
	fmt.Println(value1)

	// Encode the function setValue with a value 10 and call it
	v := big.NewInt(10)
	encodedSetValue, err := parsedABI.Pack("setValue", v)
	if err != nil {
		log.Fatal(err)
	}
	setValueTxn := jrpc.TxnArg{
		From: devAcc,
		To:   receipt.ContractAddress,
		Data: fmt.Sprintf("0x%X", encodedSetValue),
		Gas:  "0x800000",
	}
	setValueHex, err := client.SendTransaction(context.TODO(), 1, setValueTxn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(setValueHex)

	// Sleep to allow node to update the chain
	time.Sleep(1 * time.Second)

	// Get the value after calling setValue
	getValueCall2, err := client.Call(context.TODO(), 1, getValTxn, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getValueCall2)
	value2 := new(big.Int)
	value2.SetString(getValueCall2[2:], 16)
	fmt.Println(value2)

}

func getDevRandomAcct(ctx context.Context, client jrpc.Client) (string, error) {
	accts, err := client.Accounts(ctx, 1)
	if err != nil {
		return "", err
	}
	devAcct := accts[0]

	bal, err := client.GetBalance(ctx, 1, devAcct, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bal)

	return devAcct, nil
}
