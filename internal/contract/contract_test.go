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

package contract

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func Example_signTransaction() {
	privKey, err := crypto.HexToECDSA("4c0883a69102937d6231471b5dbb6204fe512961708279a8911b138d9d808759")
	if err != nil {
		fmt.Println(err)
	}
	chainID := int64(1335)
	nonce := uint64(1)
	gasTip := big.NewInt(1)
	gasPrice := big.NewInt(1)
	gasLimit := uint64(10)
	txn := createContractEIP1559Txn(chainID, nonce, gasTip, gasPrice, gasLimit, []byte("0x123a"))
	signedTxn, err := signTransaction(txn, uint64(chainID), privKey)
	if err != nil {
		fmt.Println(err)
	}

	b, err := signedTxn.MarshalBinary()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("0x%s", hex.EncodeToString(b))

	fmt.Printf("ChainID: %v\n", signedTxn.ChainId())
	fmt.Printf("Gas: %v\n", signedTxn.Gas())
	fmt.Printf("GasFeeCap: %v\n", signedTxn.GasFeeCap())
	fmt.Printf("GasTipCap: %v\n", signedTxn.GasTipCap())
	fmt.Printf("Nonce: %v\n", signedTxn.Nonce())
	v, r, s := signedTxn.RawSignatureValues()
	fmt.Printf("V: %v\nR: %v\nS: %v\n", v, r, s)

	// Output:
	// 0x02f8548205370101020a808086307831323361c080a05df2be2b626d3384b4eb9586399547cb623ad73b5d38a2d8c39562d0600a3dd7a05e9e382e8a6550b80eac9d28bf9b3766c70c0f369d24afede89347cb94f3b1bfChainID: 1335
	// Gas: 10
	// GasFeeCap: 2
	// GasTipCap: 1
	// Nonce: 1
	// V: 0
	// R: 42493984409369267390634873737391917718079187906863432014006479694048185171415
	// S: 42796957355589767142391483339700377686928456514032568249873161001150488752575
}
