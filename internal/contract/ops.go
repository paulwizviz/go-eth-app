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
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	// ErrExtractBinContent error extracting contract BIN content
	ErrExtractBinContent = errors.New("unable to extract contract bin content")
	// ErrExtractABIContent error extracting contract ABI conent
	ErrExtractABIContent = errors.New("unable to extract contract ABI content")
	// ErrSignTxn error signing transaction
	ErrSignTxn = errors.New("unable to sign transaction")
	// ErrUnableToSendTxn error sending transaction
	ErrUnableToSendTxn = errors.New("unable to send txn")
	// ErrUnableToEncodeConstructorArg error encoding constructor arg
	ErrUnableToEncodeConstructorArg = errors.New("unable to encode constructor arg")
	// ErrCreateParseABI error generating parsed ABI
	ErrCreateParseABI = errors.New("unable to create parsed ABI")
)

// ExtractContentBin extract the content of bin file
func ExtractContentBin(binFile string) (string, error) {
	return extractContractBin(binFile)
}

func extractContractBin(binFile string) (string, error) {

	// TODO: missing check of content format

	data, err := os.ReadFile(binFile)
	if err != nil {
		return "", fmt.Errorf("%w-%v", ErrExtractBinContent, err)
	}
	content := fmt.Sprintf("0x%v", string(data))
	return content, nil
}

// ExtractContractABI extract content of ABI file
func ExtractContractABI(abiFile string) (string, error) {
	return extractContractABI(abiFile)
}

func extractContractABI(abiFile string) (string, error) {

	// TODO: missing check of content

	content, err := os.ReadFile(abiFile)
	if err != nil {
		return "", fmt.Errorf("%w-%v", ErrExtractABIContent, err)
	}
	return string(content), err
}

// CreateParsedABI a wrapper to create Geth abi.ABI
func CreateParsedABI(contractABI string) (abi.ABI, error) {
	return createParsedABI(contractABI)
}

func createParsedABI(contractABI string) (abi.ABI, error) {
	parseABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("%w-%v", ErrCreateParseABI, err)
	}
	return parseABI, nil
}

// CreateContractEIP1559Txn instantiate a Dynanic Fee Transaction type for contract creation
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
