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
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
)

var (
	// ErrExtractContent error extracting contract content
	ErrExtractContent = errors.New("unable to extract contract")
	// ErrSignTxn error signing transaction
	ErrSignTxn = errors.New("unable to sign transaction")
	// ErrUnableToSendTxn error sending transaction
	ErrUnableToSendTxn = errors.New("unable to send txn")
)

// ExtractContent exteact the content of bin file
func ExtractContent(binFile string) (string, error) {
	return extractContractBin(binFile)
}

func extractContractBin(binFile string) (string, error) {
	data, err := os.ReadFile(binFile)
	if err != nil {
		return "", fmt.Errorf("%w-%v", ErrExtractContent, err)
	}
	content := fmt.Sprintf("0x%v", string(data))
	return content, nil
}

// createContractEIP1559Txn instantiate a Dynanic Fee Transaction type for contract
// creation
func CreateContractEIP1559Txn(chainID int64, nonce uint64, gasTip *big.Int, gasPrice *big.Int, gasLimit uint64, contractBin []byte) *types.Transaction {
	return createContractEIP1559Txn(chainID, nonce, gasTip, gasPrice, gasLimit, contractBin)
}

func createContractEIP1559Txn(chainID int64, nonce uint64, gasTip *big.Int, gasPrice *big.Int, gasLimit uint64, contractBin []byte) *types.Transaction {
	txData := types.DynamicFeeTx{
		ChainID:   big.NewInt(chainID), // Chain ID for the network
		Nonce:     nonce,
		GasTipCap: gasTip,                             // Tip
		GasFeeCap: new(big.Int).Add(gasPrice, gasTip), // Base fee + tip
		Gas:       gasLimit,
		To:        nil,           // `To` is nil for contract deployment
		Value:     big.NewInt(0), // Value sent with the transaction
		Data:      contractBin,   // Contract bytecode
	}
	return types.NewTx(&txData)
}

// SignTransaction is an operation to sign a transactions
func SignTransaction(txn *types.Transaction, chainID uint64, privkey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return signTransaction(txn, chainID, privkey)
}

func signTransaction(txn *types.Transaction, chainID uint64, privkey *ecdsa.PrivateKey) (*types.Transaction, error) {
	signedTx, err := types.SignTx(txn, types.LatestSignerForChainID(big.NewInt(int64(chainID))), privkey)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrSignTxn, err)
	}
	return signedTx, nil
}
