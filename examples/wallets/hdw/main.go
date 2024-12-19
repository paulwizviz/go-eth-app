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

// This example demonstrate operations related to a Hierarchical Deterministic Wallet.

package main

import (
	"fmt"
	"log"

	"github.com/paulwizviz/go-eth-app/internal/wallet"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Create Mnemonic
	m, err := wallet.Mnemonic(wallet.Mnemonic128)
	if err != nil {
		log.Fatal(err)
	}

	// Create master key with a combination of Mnemonic and passphrase
	passphrase := "hello"
	mkey, err := wallet.MasterHDKey(m, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type: %[1]T Value: %[1]v\n", mkey.Key)

	// Generate first child key
	childKey, err := mkey.NewChildKey(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(childKey)

	// Recover key from Mnemonic and passphrase
	mkey1, err := wallet.MasterHDKey(m, passphrase)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %[1]T Value: %[1]v\n", mkey1.Key)

	mkeyHex := hexutil.Encode(mkey.Key)
	mkey1Hex := hexutil.Encode(mkey1.Key)
	fmt.Println(mkeyHex, mkey1Hex)

	// Convert private key in bytes to ECDSA
	privkey, err := crypto.ToECDSA(mkey.Key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T", privkey)

}
