package jrpc

import (
	"bytes"
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
	// ErrMarshalRequest error marshaling JSON-RPC request
	ErrMarshalRequest = errors.New("marshal request error")
	// ErrSendingRequest error posting JSON-RPC request
	ErrSendingRequest = errors.New("sending request error")
	// ErrUmarshalResponse error unmarshaling JSON-RPC response
	ErrUmarshalResponse = errors.New("unmarshal respond error")
	// ErrMismatchResponse error mismatch JSON-RPC request
	ErrMismatchResponse = errors.New("mismatch response error")
)

// ErrUnmarshalBlockNumber error unmarshaling block number from JSON-RPC response
var ErrUnmarshalBlockNumber = errors.New("unmarshal block number error")

// BlockNumber returns the block number of the latest block
func BlockNumber(url string, id uint) (*big.Int, error) {
	return blockNumber(url, id)
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

// ErrUnmarshalBlock error unmarsaling block data from response
var ErrUnmarshalBlock = errors.New("unmarshal block error")

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
func GetBlockByNumber(url string, id uint, block string, hydrated bool) (Block, error) {
	return getBlockByNumber(url, id, block, hydrated)
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

// ErrUnmarshalBalance error unmarshling balance from JSON-RPC response
var ErrUnmarshalBalance = errors.New("balance error")

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
func GetBalance(url string, id uint, address string, block string) (string, error) {
	return getBalance(url, id, address, block)
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

type AccessListArg struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

// ErrTransformTxnArg error transforming TxnArg to map[string]any
var ErrTransformTxnArg = errors.New("transform txn argument error")

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

// ErrUnmarshalTxnHash error unmarshaling transaction hash from JSON-RPC response
var ErrUnmarshalTxnHash = errors.New("transaction hash error")

// SendTransaction returns a hash of the transaction
//
// Arguments:
//
//	url - to an ethereum json rpc endpoint
//	id - an identifier to match request and response
//	txn - transaction argument in object form
func SendTransaction(url string, id uint, txn TxnArg) (string, error) {
	// Convert struct to map[string]any
	m, err := transformTxnArg(txn)
	if err != nil {
		return "", err
	}
	return sendTransaction(url, id, m)
}

func sendTransaction(url string, id uint, txn map[string]any) (string, error) {

	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_sendTransaction",
		Params:  []any{txn},
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

	// Unmarshal txnHash
	var txnHash string
	if err := json.Unmarshal(rpcResp.Result, &txnHash); err != nil {
		return "", fmt.Errorf("%w-%v", ErrUnmarshalTxnHash, err)
	}
	return txnHash, nil
}
