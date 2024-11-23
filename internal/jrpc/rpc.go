package jrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"paulwizviz/go-eth-app/internal/eth"
)

var (
	ErrMarshalRequest       = errors.New("marshal request error")
	ErrUmarshalResponse     = errors.New("unmarshal respond error")
	ErrUnmarshalBlock       = errors.New("unmarshal block error")
	ErrUnmarshalBlockNumber = errors.New("unmarshal block number error")
	ErrSendingRequest       = errors.New("sending request error")
	ErrMismatchResponse     = errors.New("mismatch response error")
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

// BlockNumber returns the block number of the latest block
func BlockNumber(url string, id uint) (*big.Int, error) {

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

// GetBlockByNumber retur a block type
//
// Argment:
//
//	  url - to an ethereum json rpc endpoint
//	  id - an identifier to match request and response
//		 option - Block number or Block tag
//		 hydrated - true or false
//
//		Block number:
//		   ^0x([1-9a-f]+[0-9a-f]*|0)$
//		Block tag:
//		   `earliest`: The lowest numbered block the client has available;
//		   `finalized`: The most recent crypto-economically secure block, cannot be re-orged outside of manual intervention driven by community coordination;
//		   `safe`: The most recent block that is safe from re-orgs under honest majority and certain synchronicity assumptions;
//		   `latest`: The most recent block in the canonical chain observed by the client, this block may be re-orged out of the canonical chain even
//		             under healthy/normal conditions;
//		   `pending`: A sample next block built by the client on top of `latest` and containing the set of transactions usually taken from local mempool.
//		              Before the merge transition is finalized, any call querying for `finalized` or `safe` block MUST be responded to with
//		             `-39001: Unknown block` error
func GetBlockByNumber(url string, id uint, option string, hydrated bool) (eth.Block, error) {
	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_getBlockByNumber",
		Params:  []any{option, true}, // 'true' to get full transaction objects
		ID:      id,
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return eth.Block{}, fmt.Errorf("%w-%v", ErrMarshalRequest, err)
	}

	// Send the request
	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return eth.Block{}, fmt.Errorf("%w-%v", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	// Decode the response
	var rpcResp response
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return eth.Block{}, fmt.Errorf("%w-%v", ErrUmarshalResponse, err)
	}

	if req.ID != rpcResp.ID {
		return eth.Block{}, ErrMismatchResponse
	}

	// Unmarshal the block data (including transactions)
	var blk eth.Block
	if err := json.Unmarshal(rpcResp.Result, &blk); err != nil {
		return eth.Block{}, fmt.Errorf("%w-%v", ErrUnmarshalBlock, err)
	}

	return blk, nil
}