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
	"context"
	"encoding/json"
	"fmt"
	"time"
)

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
