package eth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
)

const (
	rpcVersion          = "2.0"
	methodBlockNumber   = "eth_blockNumber"
	methodBlockByNumber = "eth_getBlockByNumber"
)

var (
	errMarshalRequest       = errors.New("marshal request error")
	errUmarshalResponse     = errors.New("unmarshal respond error")
	errUnmarshalBlock       = errors.New("unmarshal block error")
	errUnmarshalBlockNumber = errors.New("unmarshal block number error")
	errSendingRequest       = errors.New("sending request error")
)

type request struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	ID      int    `json:"id"`
}

type response struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
}

func currentBlock(url string) (*big.Int, error) {

	req := request{
		JsonRPC: rpcVersion,
		Method:  methodBlockNumber,
		Params:  []interface{}{},
		ID:      1, // In production this would be replaced by Nounce
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", errMarshalRequest, err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("%w-%v", errSendingRequest, err)
	}
	defer resp.Body.Close()

	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, fmt.Errorf("%w-%v", errUmarshalResponse, err)
	}

	// Convert the block number from hex to decimal
	var blockNumber string
	if err := json.Unmarshal(rpcResp.Result, &blockNumber); err != nil {
		return nil, fmt.Errorf("%w-%v", errUnmarshalBlockNumber, err)
	}

	// Convert to integer
	blockNumberBig := new(big.Int)
	blockNumberBig.SetString(blockNumber[2:], 16) // Remove Ox prefix

	// Print the block number (hexadecimal format)
	return blockNumberBig, nil
}

func getBlockTransactions(url string, blockNumber *big.Int) ([]Transaction, error) {
	// Create the request body for eth_getBlockByNumber
	hexBlockNumber := fmt.Sprintf("0x%x", blockNumber) // Convert block number to hex format
	req := request{
		JsonRPC: rpcVersion,
		Method:  methodBlockByNumber,
		Params:  []interface{}{hexBlockNumber, true}, // 'true' to get full transaction objects
		ID:      1,
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", errMarshalRequest, err)
	}

	// Send the request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("%w-%v", errSendingRequest, err)
	}
	defer resp.Body.Close()

	// Decode the response
	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, fmt.Errorf("%w-%v", errUmarshalResponse, err)
	}

	// Unmarshal the block data (including transactions)
	var block Block
	if err := json.Unmarshal(rpcResp.Result, &block); err != nil {
		return nil, fmt.Errorf("%w-%v", errUnmarshalBlock, err)
	}

	return block.Transactions, nil
}
