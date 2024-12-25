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

// This example uses Go Ethereum ethclient package
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to an Ethereum node
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Fetch the latest block
	block, err := client.BlockByNumber(context.Background(), nil) // `nil` for the latest block
	if err != nil {
		log.Fatalf("Failed to retrieve the latest block: %v", err)
	}

	// Display block details
	fmt.Printf("Block Number: %d\n", block.NumberU64())
	fmt.Printf("Block Hash: %s\n", block.Hash().Hex())
	fmt.Printf("Parent Hash: %s\n", block.ParentHash().Hex())
	fmt.Printf("Miner: %s\n", block.Coinbase().Hex())
	fmt.Printf("Gas Used: %d\n", block.GasUsed())
	fmt.Printf("Transactions Count: %d\n", len(block.Transactions()))

	// Iterate through transactions
	for _, tx := range block.Transactions() {
		fmt.Printf("Tx Hash: %s\n", tx.Hash().Hex())
	}
}
