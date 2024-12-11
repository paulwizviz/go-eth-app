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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
)

const (
	rpcVersion  = "2.0"
	contentType = "application/json"
)

type request struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	ID      uint   `json:"id"`
}

type response struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      uint            `json:"id"`
	Result  json.RawMessage `json:"result"`
}

const (

	// BlockTagEARLEST (`earliest`): The lowest numbered block the client has available;
	BlockTagEARLEST = "earliest"

	// BlockTagFINALIZED (`finalized`): The most recent crypto-economically secure block,
	//                                  cannot be re-orged outside of manual intervention
	//                                  driven by community coordination;
	BlockTagFinalized = "finalized"

	// BlockTagSAFE	(`safe`): The most recent block that is safe from re-orgs under honest
	//                        majority and certain synchronicity assumptions;
	BlockTagSAFE = "safe"

	// BlockTagLATEST (`latest`): The most recent block in the canonical chain observed by
	//                            the client, this block may be re-orged out of the canonical
	//                            chain even under healthy/normal conditions;
	BlockTagLATEST = "latest"

	//	BlockTagPENDING (`pending`): A sample next block built by the client on top of `latest`
	//                               and containing the set of transactions usually taken from
	//                               local mempool. Before the merge transition is finalized,
	//                               any call querying for `finalized` or `safe` block MUST be
	//                               responded to with  d`-39001: Unknown block` error
	BlockTagPENDING = "pending"
)

var (
	// ErrContextCancelRPC context initated a cancel
	ErrContextCancelRPC = errors.New("context cancel called")
	// ErrMarshalRequest error marshaling JSON-RPC request
	ErrMarshalRequest = errors.New("marshal request")
	// ErrSendingRequest error posting JSON-RPC request
	ErrSendingRequest = errors.New("sending request")
	// ErrUmarshalResponse error unmarshaling JSON-RPC response
	ErrUmarshalResponse = errors.New("unmarshal respond")
	// ErrMismatchResponse error mismatch JSON-RPC request
	ErrMismatchResponse = errors.New("mismatch response")
	// ErrUnmarshalAccounts error unmarshaling accout list
	ErrUnmarshalAccounts = errors.New("umarshal accounts")
	// ErrUnmarshalBlockNumber error unmarshaling block number from JSON-RPC response
	ErrUnmarshalBlockNumber = errors.New("unmarshal block number")
	// ErrUnmarshalBlock error unmarsaling block data from response
	ErrUnmarshalBlock = errors.New("unmarshal block")
	// ErrUnmarshalBalance error unmarshling balance from JSON-RPC response
	ErrUnmarshalBalance = errors.New("umarshal balance")
	// ErrTransformTxnArg error transforming TxnArg to map[string]any
	ErrTransformTxnArg = errors.New("transform txn argument")
	// ErrUnmarshalTxnHash error unmarshaling transaction hash from JSON-RPC response
	ErrUnmarshalTxnHash = errors.New("unmarshal transaction hash")
	// ErrUnmarshalNetworkID error unmarshaling networkID
	ErrUnmarshalNetworkID = errors.New("unmarhsal network id")
	// ErrUnmarshalGasPrice error unmarshaling gas price
	ErrUnmarshalGasPrice = errors.New("unmarshal gas price")
	// ErrUnmarshalTxnCount error unmarshaling txn count
	ErrUnmarshalTxnCount = errors.New("unmarshal txn count")
)

// Accounts return a list of accounts owned by the client
func Accounts(ctx context.Context, url string, id uint) ([]string, error) {
	select {
	case <-ctx.Done():
		return nil, ErrContextCancelRPC
	default:
		return accounts(url, id)
	}
}

func accounts(url string, id uint) ([]string, error) {
	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_accounts",
		Params:  []any{},
		ID:      id,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}

	if req.ID != rpcResp.ID {
		return nil, ErrMismatchResponse
	}

	// Convert the block number from hex to decimal
	var accts []string
	if err := json.Unmarshal(rpcResp.Result, &accts); err != nil {
		return nil, fmt.Errorf("%w-%v", ErrUnmarshalBlockNumber, err)
	}

	return accts, nil
}

// BlockNumber returns the block number of the latest block
func BlockNumber(ctx context.Context, url string, id uint) (*big.Int, error) {
	select {
	case <-ctx.Done():
		return nil, ErrContextCancelRPC
	default:
		return blockNumber(url, id)
	}
}

func blockNumber(url string, id uint) (*big.Int, error) {

	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_blockNumber",
		Params:  []any{},
		ID:      id,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}

	if req.ID != rpcResp.ID {
		return nil, ErrMismatchResponse
	}

	// Convert the block number from hex to decimal
	var blkNum string
	if err := json.Unmarshal(rpcResp.Result, &blkNum); err != nil {
		return nil, fmt.Errorf("%w-%v", ErrUnmarshalBlockNumber, err)
	}

	// Convert to integer
	blockNumber := new(big.Int)
	blockNumber.SetString(blkNum[2:], 16) // Remove Ox prefix

	return blockNumber, nil
}

// GasPrice return suggested gas price in int64 (wei)
func GasPrice(ctx context.Context, url string, id uint) (*big.Int, error) {
	select {
	case <-ctx.Done():
		return nil, ErrContextCancelRPC
	default:
		return gasPrice(url, id)
	}
}

func gasPrice(url string, id uint) (*big.Int, error) {

	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_gasPrice",
		Params:  []any{},
		ID:      id,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}

	if req.ID != rpcResp.ID {
		return big.NewInt(-1), ErrMismatchResponse
	}

	// Convert the block number from hex to decimal
	var netID string
	if err := json.Unmarshal(rpcResp.Result, &netID); err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrUnmarshalNetworkID, err)
	}

	networkID := new(big.Int)
	networkID.SetString(netID[2:], 16)

	return networkID, nil
}

// GetBlockByNumber returns a block type
//
// Argments:
//
//	url - to an ethereum json rpc endpoint
//	id - an identifier to match request and response
//	block - Block number or Block tag
//		Block number: ^0x([1-9a-f]+[0-9a-f]*|0)$
//		Block tag: See constants
//	hydrated - true or false
func GetBlockByNumber(ctx context.Context, url string, id uint, block string, hydrated bool) (Block, error) {
	select {
	case <-ctx.Done():
		return Block{}, ErrContextCancelRPC
	default:
		return getBlockByNumber(url, id, block, hydrated)
	}
}

func getBlockByNumber(url string, id uint, block string, hydrated bool) (Block, error) {
	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_getBlockByNumber",
		Params:  []any{block, hydrated},
		ID:      id,
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return Block{}, fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}

	// Send the request
	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return Block{}, fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	// Decode the response
	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return Block{}, fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}

	if req.ID != rpcResp.ID {
		return Block{}, ErrMismatchResponse
	}

	// Unmarshal the block data (including transactions)
	var blk Block
	if err := json.Unmarshal(rpcResp.Result, &blk); err != nil {
		return Block{}, fmt.Errorf("%w-%v", ErrUnmarshalBlock, err)
	}

	return blk, nil
}

// GetBalance returns the balance for a given address and block
//
// Arguments:
//
//	url - to an ethereum json rpc endpoint
//	id - an identifier to match request and response
//	address - of the account
//	block - Block number or Block tag
//		Block number: ^0x([1-9a-f]+[0-9a-f]*|0)$
//		Block tag: See constants
func GetBalance(ctx context.Context, url string, id uint, address string, block string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ErrContextCancelRPC
	default:
		return getBalance(url, id, address, block)
	}
}

func getBalance(url string, id uint, address string, block string) (string, error) {
	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_getBalance",
		Params:  []any{address, block},
		ID:      id,
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}

	// Send the request
	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	// Decode the response
	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return "", fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}

	if req.ID != rpcResp.ID {
		return "", ErrMismatchResponse
	}

	// Unmarshal balance
	var balance string
	if err := json.Unmarshal(rpcResp.Result, &balance); err != nil {
		return "", fmt.Errorf("%w-%v", ErrUnmarshalBalance, err)
	}

	return balance, nil
}

// GetTxnCount returns the nonce in big.Int depending on status
//
// Arguments:
//
//	url - to an ethereum json rpc endpoint
//	id - an identifier to match request and response
//	address - of the account
//	block - Block number or Block tag
//		Block number: ^0x([1-9a-f]+[0-9a-f]*|0)$
//		Block tag: See constants
func GetTxnCount(ctx context.Context, url string, id uint, address string, block string) (*big.Int, error) {
	select {
	case <-ctx.Done():
		return nil, ErrContextCancelRPC
	default:
		return getTxnCount(url, id, address, block)
	}
}

func getTxnCount(url string, id uint, address string, block string) (*big.Int, error) {
	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_getBalance",
		Params:  []any{address, block},
		ID:      id,
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}

	// Send the request
	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	// Decode the response
	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}

	if req.ID != rpcResp.ID {
		return big.NewInt(-1), ErrMismatchResponse
	}

	// Unmarshal count
	var count string
	if err := json.Unmarshal(rpcResp.Result, &count); err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrUnmarshalTxnCount, err)
	}

	ct := new(big.Int)
	ct.SetString(count[2:], 16)
	return ct, nil
}

// NetworkID returns the ID in int64
func NetworkID(ctx context.Context, url string, id uint) (*big.Int, error) {
	select {
	case <-ctx.Done():
		return nil, ErrContextCancelRPC
	default:
		return networkID(url, id)
	}
}

func networkID(url string, id uint) (*big.Int, error) {

	req := request{
		JsonRPC: rpcVersion,
		Method:  "net_version",
		Params:  []any{},
		ID:      id,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}

	if req.ID != rpcResp.ID {
		return big.NewInt(-1), ErrMismatchResponse
	}

	// Convert the block number from hex to decimal
	var networkID string
	if err := json.Unmarshal(rpcResp.Result, &networkID); err != nil {
		return big.NewInt(-1), fmt.Errorf("%w-%v", ErrUnmarshalNetworkID, err)
	}

	netID := new(big.Int)
	netID.SetString(networkID, 10)
	return netID, nil
}

type AccessListArg struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

// TxnArg argument for send transaction call
type TxnArg struct {
	Type                 string          `json:"type,omitempty"`
	Nonce                string          `json:"nonce,omitempty"`
	To                   string          `json:"to,omitempty"`
	From                 string          `json:"from,omitempty"`
	Gas                  string          `json:"gas,omitempty"`
	Value                string          `json:"value,omitempty"`
	Input                string          `json:"input,omitempty"`
	GasPrice             string          `json:"gasPrice,omitempty"`
	MaxPriorityFeePerGas string          `json:"maxPriorityFeePerGas,omitempty"`
	MaxFreePerGas        string          `json:"maxFeePerGas,omitempty"`
	MaxFeePerBlobGas     string          `json:"maxFeePerBlobGas,omitempty"`
	AccessList           []AccessListArg `json:"accessList,omitempty"`
	BlobVersionedHashes  []string        `json:"blobVersionedHashes,omitempty"`
	Blobs                []string        `json:"blobs,omitempty"`
	ChainID              string          `json:"chainId,omitempty"`
}

func transformTxnArg(txn TxnArg) (map[string]any, error) {

	b, err := json.Marshal(txn)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrTransformTxnArg, err)
	}

	var m map[string]any
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrTransformTxnArg, err)
	}

	return m, nil
}

// SendTransaction returns a hash of the transaction.
// NOTE: Use this for cases where the private key is stored on the node.
//
// Arguments:
//
//	url - to an ethereum json rpc endpoint
//	id - an identifier to match request and response
//	txn - transaction of type TxnArg
func SendTransaction(ctx context.Context, url string, id uint, txn TxnArg) (string, error) {
	select {
	case <-ctx.Done():
		return "", ErrContextCancelRPC
	default:
		// Convert struct to map[string]any
		m, err := transformTxnArg(txn)
		if err != nil {
			return "", err
		}
		return sendTransaction(url, id, m)
	}
}

// SendRawTransaction returns a hash of the transaction
// NOTE: Use this for cases where you sign transaction externally and
//
//	passed the signed the transaction.
//
// Arguments:
//
//	url - to an ethereum json rpc endpoint
//	id - an identifier to match request and response
//	txn - signed transaction hex in string
func SendRawTransaction(ctx context.Context, url string, id uint, txn string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ErrContextCancelRPC
	default:
		return sendRawTransaction(url, id, txn)
	}
}
