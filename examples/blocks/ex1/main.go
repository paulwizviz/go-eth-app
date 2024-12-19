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

// This example demonstrate operations to explore blocks in a public network
// using the internal implementation of JSON-RPC.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/paulwizviz/go-eth-app/internal/jrpc"
)

func main() {

	// Instantiate a default client of internal JSON-RPC package
	client := jrpc.NewDefaultClient("https://ethereum-rpc.publicnode.com")

	// Get the most recent block
	number, err := client.BlockNumber(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}

	// Get detail of the block
	blknum := fmt.Sprintf("0x%x", number)
	blk, err := client.GetBlockByNumber(context.TODO(), 1, blknum, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%[1]T", blk.Transactions)
	switch v := blk.Transactions.(type) {
	case []jrpc.TxnEIP1559:
		for _, tx := range v {
			fmt.Println(tx.BlockNumber)
		}
	case []jrpc.TxnPreEIP2718:
		for _, tx := range v {
			fmt.Println(tx.BlockNumber)
		}
	case []jrpc.TxnEIP2718:
		for _, tx := range v {
			fmt.Println(tx.BlockNumber)
		}
	case []jrpc.TxnEIP2930:
		for _, tx := range v {
			fmt.Println(tx.Data)
		}
	case []jrpc.TxnEIP4844:
		for _, tx := range v {
			fmt.Println(tx.BlockNumber)
		}
	}
}
