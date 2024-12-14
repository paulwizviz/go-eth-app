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
	"time"
)

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
	// ErrFormRequest error forming request error
	ErrFormRequest = errors.New("unable to form request")
	// ErrSendingRequest error posting JSON-RPC request
	ErrSendingRequest = errors.New("sending request")
	// ErrUmarshalResponse error unmarshaling JSON-RPC response
	ErrUmarshalResponse = errors.New("unmarshal respond")
	// ErrMismatchResponse error mismatch JSON-RPC request
	ErrMismatchResponse = errors.New("mismatch response")
	// ErrResponse error from the JSON-RPC server
	ErrResponse = errors.New("response error")
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
	// ErrUnmarshalTxnHash error unable to get count
	ErrUnmarshalTxnHash = errors.New("unable to get transaction count")
	// ErrUnmarshalNetworkID error unmarshaling networkID
	ErrUnmarshalNetworkID = errors.New("unmarhsal network id")
	// ErrUnmarshalGasPrice error unmarshaling gas price
	ErrUnmarshalGasPrice = errors.New("unmarshal gas price")
	// ErrUnmarshalCallHash error unmarshaling call hash
	ErrUnmarshalCallHash = errors.New("unmarshall call hash")
	// ErrTxnCount error unmarshaling txn count
	ErrTxnCount = errors.New("unable to get txn count")
	// ErrUnmarshalTxnReceipt error unmarshaling receipt
	ErrUnmarshalTxnReceipt = errors.New("unable to unmarshal txn receipt")
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

func generateRequestBody(reqID uint, method string, params []any) ([]byte, error) {
	req := request{
		JsonRPC: rpcVersion,
		ID:      reqID,
		Method:  method,
		Params:  params,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}
	return b, nil
}

type respError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type response struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      uint            `json:"id"`
	Result  json.RawMessage `json:"result"`
	Err     *respError      `json:"error,omitempty"`
}

func postRPC(ctx context.Context, timeout time.Duration, url string, reqID uint, reqBody []byte) (response, error) {
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return response{}, fmt.Errorf("%w-%v", ErrFormRequest, err)
	}
	request.Header.Add("Content-Type", contentType)
	req := request.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return response{}, fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return response{}, fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}
	if reqID != rpcResp.ID {
		return response{}, ErrMismatchResponse
	}

	if rpcResp.Err != nil {
		errMsg := fmt.Sprintf("Error code: %v message: %v", rpcResp.Err.Code, rpcResp.Err.Message)
		return response{}, fmt.Errorf("%w-%v", ErrResponse, errMsg)
	}
	return rpcResp, nil
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
	Data                 string          `json:"data,omitempty"`
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

// TxnReceipt represents receipt from transaction hash
type TxnReceipt struct {
	Type              string   `json:"type,omitempty"`
	TransactionHash   string   `json:"transactionHash,omitempty"`
	BlockHash         string   `json:"blockHash,omitempty"`
	BlockNumber       string   `json:"blockNumber,omitempty"`
	From              string   `json:"from,omitempty"`
	To                string   `json:"to,omitempty"`
	CumulativeGasUsed string   `json:"cumulativeGasUsed,omitempty"`
	ContractAddress   string   `json:"contractAddress,omitempty"`
	Logs              []string `json:"logs,omitempty"`
	LogsBloom         string   `json:"logsBloom,omitempty"`
	Root              string   `json:"root,omitempty"`
	Status            string   `json:"status,omitempty"`
	EffectiveGasPrice string   `json:"effectiveGasPrice,omitempty"`
	BlobGasPrice      string   `json:"blobGasPrice,omitempty"`
}

// Client represent a http client
type Client interface {
	// Accounts return a list of accounts owned by the reference node
	Accounts(ctx context.Context, reqID uint) ([]string, error)
	// BlockNumber returns the number of the most recent block
	BlockNumber(ctx context.Context, reqID uint) (*big.Int, error)
	// Call executes a new message call immediately without creating a transaction
	Call(ctx context.Context, reqID uint, txn TxnArg, block string) (string, error)
	// GasPrice return suggested gas price in int64 (wei)
	GasPrice(ctx context.Context, reqID uint) (*big.Int, error)
	// GetBlockByNumber returns a block type
	//
	// Argments:
	//
	//	reqID - an identifier to match request and response
	//	block - Block number or Block tag
	//		Block number: ^0x([1-9a-f]+[0-9a-f]*|0)$
	//		Block tag: See constants
	//	hydrated - true or false
	GetBlockByNumber(ctx context.Context, reqID uint, block string, hydrated bool) (Block, error)
	// GetBalance returns the balance for a given address and block
	//
	// Arguments:
	//
	//	reqID - an identifier to match request and response
	//	address - of the account
	//	block - Block number or Block tag
	//		Block number: ^0x([1-9a-f]+[0-9a-f]*|0)$
	//		Block tag: See constants
	GetBalance(ctx context.Context, reqID uint, address string, block string) (string, error)
	// GetTxnCount returns the nonce in big.Int depending on status.
	//
	// Arguments:
	//
	//	reqID - an identifier to match request and response
	//	address - of the account
	//	block - Block number or Block tag
	//		Block number: ^0x([1-9a-f]+[0-9a-f]*|0)$
	//		Block tag: See constants
	GetTxnCount(ctx context.Context, reqID uint, address string, block string) (*big.Int, error)
	// GetTxnReceipt return receipt for a given txnHash
	GetTxnReceipt(ctx context.Context, reqID uint, txnHash string) (TxnReceipt, error)
	// NetworkID returns the ID in int64
	NetworkID(ctx context.Context, reqID uint) (*big.Int, error)
	// SendTransaction returns a hash of the transaction.
	// NOTE: Use this for cases where the private key is stored on the node.
	//
	// Arguments:
	//
	//	reqID - an identifier to match request and response
	//	txn - transaction of type TxnArg
	SendTransaction(ctx context.Context, reqID uint, txn TxnArg) (string, error)
	// SendRawTransaction returns a hash of the transaction
	// NOTE: Use this for cases where you sign transaction externally and
	//	passed the signed the transaction.
	//
	// Arguments:
	//
	//	reqID - an identifier to match request and response
	//	txn - signed transaction hex in string
	SendRawTransaction(ctx context.Context, reqID uint, txn string) (string, error)
}

type client struct {
	timeout time.Duration
	url     string
}

func (c client) Accounts(ctx context.Context, reqID uint) ([]string, error) {
	return accounts(ctx, c.timeout, c.url, reqID)
}

func accounts(ctx context.Context, timeout time.Duration, url string, reqID uint) ([]string, error) {

	reqBody, err := generateRequestBody(reqID, "eth_accounts", []any{})
	if err != nil {
		return nil, err
	}

	response, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return nil, err
	}

	// Convert the block number from hex to decimal
	var accts []string
	if err := json.Unmarshal(response.Result, &accts); err != nil {
		return nil, fmt.Errorf("%w-%v", ErrUnmarshalBlockNumber, err)
	}

	return accts, nil
}

func (c client) BlockNumber(ctx context.Context, reqID uint) (*big.Int, error) {
	return blockNumber(ctx, c.timeout, c.url, reqID)
}

func blockNumber(ctx context.Context, timeout time.Duration, url string, reqID uint) (*big.Int, error) {

	reqBody, err := generateRequestBody(reqID, "eth_blockNumber", []any{})
	if err != nil {
		return nil, err
	}

	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return nil, err
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

func (c client) Call(ctx context.Context, reqID uint, txn TxnArg, block string) (string, error) {
	// Convert struct to map[string]any
	m, err := transformTxnArg(txn)
	if err != nil {
		return "", err
	}
	return call(ctx, c.timeout, c.url, reqID, m, block)
}

func call(ctx context.Context, timeout time.Duration, url string, reqID uint, txn map[string]any, block string) (string, error) {

	reqBody, err := generateRequestBody(reqID, "eth_call", []any{txn, block})
	if err != nil {
		return "", err
	}

	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return "", err
	}

	// Unmarshal callhash
	var callHash string
	if err := json.Unmarshal(rpcResp.Result, &callHash); err != nil {
		return "", fmt.Errorf("%w-%v", ErrUnmarshalCallHash, err)
	}

	return callHash, nil
}

func (c client) GasPrice(ctx context.Context, reqID uint) (*big.Int, error) {
	return gasPrice(ctx, c.timeout, c.url, reqID)
}

func gasPrice(ctx context.Context, timeout time.Duration, url string, reqID uint) (*big.Int, error) {
	reqBody, err := generateRequestBody(reqID, "eth_gasPrice", []any{})
	if err != nil {
		return nil, err
	}
	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return nil, err
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

func (c client) GetBlockByNumber(ctx context.Context, reqID uint, block string, hydrated bool) (Block, error) {
	return getBlockByNumber(ctx, c.timeout, c.url, reqID, block, hydrated)
}

func getBlockByNumber(ctx context.Context, timeout time.Duration, url string, reqID uint, block string, hydrated bool) (Block, error) {

	reqBody, err := generateRequestBody(reqID, "eth_getBlockByNumber", []any{block, hydrated})
	if err != nil {
		return Block{}, err
	}

	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return Block{}, err
	}

	// Unmarshal the block data (including transactions)
	var blk Block
	if err := json.Unmarshal(rpcResp.Result, &blk); err != nil {
		return Block{}, fmt.Errorf("%w-%v", ErrUnmarshalBlock, err)
	}

	return blk, nil
}

func (c client) GetBalance(ctx context.Context, reqID uint, address string, block string) (string, error) {
	return getBalance(ctx, c.timeout, c.url, reqID, address, block)
}

func getBalance(ctx context.Context, timeout time.Duration, url string, reqID uint, address string, block string) (string, error) {

	reqBody, err := generateRequestBody(reqID, "eth_getBalance", []any{address, block})
	if err != nil {
		return "", err
	}

	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return "", err
	}

	// Unmarshal balance
	var balance string
	if err := json.Unmarshal(rpcResp.Result, &balance); err != nil {
		return "", fmt.Errorf("%w-%v", ErrUnmarshalBalance, err)
	}

	return balance, nil
}

func (c client) GetTxnCount(ctx context.Context, reqID uint, address string, block string) (*big.Int, error) {
	return getTxnCount(ctx, c.timeout, c.url, reqID, address, block)
}

func getTxnCount(ctx context.Context, timeout time.Duration, url string, reqID uint, address string, block string) (*big.Int, error) {
	count, err := getBalance(ctx, timeout, url, reqID, address, block)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrTxnCount, err)
	}
	ct := new(big.Int)
	ct.SetString(count[2:], 16)
	return ct, nil
}

func (c client) GetTxnReceipt(ctx context.Context, reqID uint, txnHash string) (TxnReceipt, error) {
	return getTxnReceipt(ctx, c.timeout, c.url, reqID, txnHash)
}

func getTxnReceipt(ctx context.Context, timeout time.Duration, url string, reqID uint, txnHash string) (TxnReceipt, error) {
	reqBody, err := generateRequestBody(reqID, "eth_getTransactionReceipt", []any{txnHash})
	if err != nil {
		return TxnReceipt{}, err
	}

	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return TxnReceipt{}, err
	}

	// Unmarshal transaction receipt
	var receipt TxnReceipt
	if err := json.Unmarshal(rpcResp.Result, &receipt); err != nil {
		return TxnReceipt{}, fmt.Errorf("%w-%v", ErrUnmarshalTxnReceipt, err)
	}

	return receipt, nil
}

func (c client) NetworkID(ctx context.Context, reqID uint) (*big.Int, error) {
	return networkID(ctx, c.timeout, c.url, reqID)
}

func networkID(ctx context.Context, timeout time.Duration, url string, reqID uint) (*big.Int, error) {

	reqBody, err := generateRequestBody(reqID, "net_version", []any{})
	if err != nil {
		return nil, err
	}

	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return nil, err
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

func (c client) SendTransaction(ctx context.Context, reqID uint, txn TxnArg) (string, error) {
	// Convert struct to map[string]any
	m, err := transformTxnArg(txn)
	if err != nil {
		return "", err
	}
	return sendTransaction(ctx, c.timeout, c.url, reqID, m)
}

func sendTransaction(ctx context.Context, timeout time.Duration, url string, reqID uint, txn map[string]any) (string, error) {

	reqBody, err := generateRequestBody(reqID, "eth_sendTransaction", []any{txn})
	if err != nil {
		return "", err
	}

	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return "", err
	}

	// Unmarshal txnHash
	var txnHash string
	if err := json.Unmarshal(rpcResp.Result, &txnHash); err != nil {
		return "", fmt.Errorf("%w-%v", ErrUnmarshalTxnHash, err)
	}
	return txnHash, nil
}

func (c client) SendRawTransaction(ctx context.Context, reqID uint, txn string) (string, error) {
	return sendRawTransaction(ctx, c.timeout, c.url, reqID, txn)
}

func sendRawTransaction(ctx context.Context, timeout time.Duration, url string, reqID uint, txn string) (string, error) {

	reqBody, err := generateRequestBody(reqID, "eth_sendRawTransaction", []any{txn})
	if err != nil {
		return "", err
	}

	rpcResp, err := postRPC(ctx, timeout, url, reqID, reqBody)
	if err != nil {
		return "", err
	}

	// Unmarshal txnHash
	var txnHash string
	if err := json.Unmarshal(rpcResp.Result, &txnHash); err != nil {
		return "", fmt.Errorf("%w-%v", ErrUnmarshalTxnHash, err)
	}
	return txnHash, nil
}

func NewDefaultClient(url string) Client {
	return client{
		timeout: 60 * time.Second,
		url:     url,
	}
}

func NewClient(timeout time.Duration, url string) Client {
	return client{
		timeout: timeout,
		url:     url,
	}
}
