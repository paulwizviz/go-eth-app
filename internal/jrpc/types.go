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

package jrpc

import (
	"encoding/json"
)

type TxnType int

type AccessList struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

// Txn is a structure for storing transaction data
type Txn interface{}

// TxnPreEIP2718 is a representation of Pre EIP-2718 transaction
type TxnPreEIP2718 struct {
	Type             *string `json:"type"`
	Nonce            string  `json:"nonce"`
	GasPrice         string  `json:"gasPrice"`
	Gas              string  `json:"gas"`
	To               string  `json:"to"`
	Value            string  `json:"value"`
	Input            string  `json:"input"`
	V                string  `json:"v"`
	R                string  `json:"r"`
	S                string  `json:"s"`
	Hash             string  `json:"hash"`
	BlockHash        string  `json:"blockHash"`
	BlockNumber      string  `json:"blockNumber"`
	TransactionIndex string  `json:"transactionIndex"`
}

// TxnEIP2718 is a representation of EIP-2718 transaction.
// Type: 0x0
type TxnEIP2718 struct {
	Type             *string `json:"type"`
	Nonce            string  `json:"nonce"`
	GasPrice         string  `json:"gasPrice"`
	Gas              string  `json:"gas"`
	To               string  `json:"to"`
	Value            string  `json:"value"`
	Input            string  `json:"input"`
	V                string  `json:"v"`
	R                string  `json:"r"`
	S                string  `json:"s"`
	Hash             string  `json:"hash"`
	BlockHash        string  `json:"blockHash"`
	BlockNumber      string  `json:"blockNumber"`
	TransactionIndex string  `json:"transactionIndex"`
}

// TxnEIP2930 is a representation of EIP2930 transaction.
// Type: 0x1
type TxnEIP2930 struct {
	Type             *string      `json:"type"`
	ChainID          string       `json:"chainId"`
	Nonce            string       `json:"nonce"`
	GasPrice         string       `json:"gasPrice"`
	Gas              string       `json:"gas"`
	To               string       `json:"to"`
	Value            string       `json:"value"`
	Input            string       `json:"input"`
	AccessList       []AccessList `json:"accessList"`
	V                string       `json:"v"`
	R                string       `json:"r"`
	S                string       `json:"s"`
	Hash             string       `json:"hash"`
	BlockHash        string       `json:"blockHash"`
	BlockNumber      string       `json:"blockNumber"`
	TransactionIndex string       `json:"transactionIndex"`
}

// TxnEIP1559 is a representation of EIP2930 transaction.
// Type: 0x2
type TxnEIP1559 struct {
	Type                 *string `json:"type"`
	ChainID              string  `json:"chainId"`
	Nonce                string  `json:"nonce"`
	MaxPriorityFeePerGas string  `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string  `json:"maxFeePerGas"`
	Gas                  string  `json:"gas"`
	To                   string  `json:"to"`
	Value                string  `json:"value"`
	Input                string  `json:"input"`
	V                    string  `json:"v"`
	R                    string  `json:"r"`
	S                    string  `json:"s"`
	Hash                 string  `json:"hash"`
	BlockHash            string  `json:"blockHash"`
	BlockNumber          string  `json:"blockNumber"`
	TransactionIndex     string  `json:"transactionIndex"`
}

// TxnEIP4844 is a representation of EIP-4844 transaction.
// Type: 0x5
type TxnEIP4844 struct {
	Type                 *string  `json:"type"`
	ChainID              string   `json:"chainId"`
	Nonce                string   `json:"nonce"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string   `json:"maxFeePerGas"`
	Gas                  string   `json:"gas"`
	To                   string   `json:"to"`
	Value                string   `json:"value"`
	Input                string   `json:"input"`
	BlobVersonedHashes   []string `json:"blobVersionedHashes"`
	V                    string   `json:"v"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	Hash                 string   `json:"hash"`
	BlockHash            string   `json:"blockHash"`
	BlockNumber          string   `json:"blockNumber"`
	TransactionIndex     string   `json:"transactionIndex"`
}

var (
	PreEIP2718  = -1
	TypeEIP2718 = "0x0"
	TypeEIP2930 = "0x1"
	TypeEIP1559 = "0x2"
	TypeEIP4844 = "0x5"
)

// Block is a representation of a block from Ethereum
// node
type Block struct {
	Number       string `json:"number"`
	Transactions Txn    `json:"transactions"`
}

func (b *Block) UnmarshalJSON(data []byte) error {
	type Alias Block
	temp := struct {
		Transactions json.RawMessage `json:"transactions"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	var txns []string
	if err := json.Unmarshal(temp.Transactions, &txns); err == nil {
		b.Transactions = txns
		return nil
	}

	var txnPreEIP2718 []TxnPreEIP2718
	if err := json.Unmarshal(temp.Transactions, &txnPreEIP2718); err == nil && txnPreEIP2718 != nil && txnPreEIP2718[0].Type == nil {
		b.Transactions = txnPreEIP2718
		return nil
	}

	var txnEIP2718 []TxnEIP2718
	if err := json.Unmarshal(temp.Transactions, &txnEIP2718); err == nil && txnEIP2718 != nil && txnEIP2718[0].Type != nil && *txnEIP2718[0].Type == TypeEIP2718 {
		b.Transactions = txnEIP2718
		return nil
	}

	var txnEIP2930 []TxnEIP2930
	if err := json.Unmarshal(temp.Transactions, &txnEIP2930); err == nil && txnEIP2930 != nil && txnEIP2930[0].Type != nil && *txnEIP2930[0].Type == TypeEIP2930 {
		b.Transactions = txnEIP2930
		return nil
	}

	var txnEIP1559 []TxnEIP1559
	if err := json.Unmarshal(temp.Transactions, &txnEIP1559); err == nil && txnEIP1559 != nil && txnEIP1559[0].Type != nil && *txnEIP1559[0].Type == TypeEIP1559 {
		b.Transactions = txnEIP1559
		return nil
	}

	var txnEIP4844 []TxnEIP4844
	if err := json.Unmarshal(temp.Transactions, &txnEIP4844); err == nil && txnEIP4844 != nil && txnEIP4844[0].Type != nil && *txnEIP4844[0].Type == TypeEIP4844 {
		b.Transactions = txnEIP4844
		return nil
	}

	return ErrUnmarshalBlock
}
