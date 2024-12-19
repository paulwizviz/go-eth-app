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

func Example_extractContractBin() {
	content, err := extractContractBin("./testdata/HelloWorld.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content)

	// Output:
	// 0x6080604052348015600e575f5ffd5b5060405161020f38038061020f8339818101604052810190602e9190606b565b805f81905550506091565b5f5ffd5b5f819050919050565b604d81603d565b81146056575f5ffd5b50565b5f815190506065816046565b92915050565b5f60208284031215607d57607c6039565b5b5f6088848285016059565b91505092915050565b6101718061009e5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063209652551461004357806355241077146100615780636d619daa1461007d575b5f5ffd5b61004b61009b565b60405161005891906100c9565b60405180910390f35b61007b60048036038101906100769190610110565b6100a3565b005b6100856100ac565b60405161009291906100c9565b60405180910390f35b5f5f54905090565b805f8190555050565b5f5481565b5f819050919050565b6100c3816100b1565b82525050565b5f6020820190506100dc5f8301846100ba565b92915050565b5f5ffd5b6100ef816100b1565b81146100f9575f5ffd5b50565b5f8135905061010a816100e6565b92915050565b5f60208284031215610125576101246100e2565b5b5f610132848285016100fc565b9150509291505056fea26469706673582212207c606765efd9417678f0dc77a0d3b0866f663226169b3aa1718c01d0c67d8ea664736f6c634300081c0033
}

func Example_extractContractABI() {

	content, err := extractContractABI("./testdata/HelloWorld.abi")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content)

	// Output:
	// [{"inputs":[{"internalType":"uint256","name":"initialValue","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"getValue","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"newValue","type":"uint256"}],"name":"setValue","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"storedValue","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]
}

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
