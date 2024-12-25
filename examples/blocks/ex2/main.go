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

// This example demonstrates operations to explorer local node
// in Dev mode

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	// Instantiate a default client of internal JSON-RPC package
	client := jrpc.NewDefaultClient("http://localhost:8545")

	// Get the most recent block
	number, err := client.BlockNumber(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}
	blknum := hexutil.EncodeBig(number)
	fmt.Println(number, blknum)

}
