package jrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ErrUnmarshalTxnHash error unmarshaling transaction hash from JSON-RPC response
var ErrUnmarshalTxnHash = errors.New("transaction hash error")

func sendRawTransaction(url string, id uint, txn string) (string, error) {
	req := request{
		JsonRPC: rpcVersion,
		Method:  "eth_sendRawTransaction",
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
