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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	testcases := []struct {
		input []byte
		want  Txn
	}{
		{
			input: []byte(`{"number":"1","transactions":["0x1","0x2"]}`),
			want:  []string{"0x1", "0x2"},
		},
		{
			input: []byte(`{
	"number": "1",
  "transactions": [
    {
      "nonce": "0x1",
      "gasPrice": "0x3b9aca00",
      "gas": "0x5208",
      "to": "0xrecipientAddress",
      "value": "0xde0b6b3a7640000",
      "input": "0x",
      "v": "0x1c",
      "r": "0xSignatureRValue",
      "s": "0xSignatureSValue",
      "hash": "0xTransactionHash",
      "blockHash": "0xBlockHash",
      "blockNumber": "0x1b4",
      "transactionIndex": "0x0"
    }
  ]
}`),
			want: []TxnPreEIP2718{
				{
					Nonce:            "0x1",
					GasPrice:         "0x3b9aca00",
					Gas:              "0x5208",
					To:               "0xrecipientAddress",
					Value:            "0xde0b6b3a7640000",
					Input:            "0x",
					V:                "0x1c",
					R:                "0xSignatureRValue",
					S:                "0xSignatureSValue",
					Hash:             "0xTransactionHash",
					BlockHash:        "0xBlockHash",
					BlockNumber:      "0x1b4",
					TransactionIndex: "0x0",
				},
			},
		},
		{
			input: []byte(`{
				"number": "1",
	  "transactions": [
		{
		  "type": "0x0",
		  "nonce": "0x1",
		  "gasPrice": "0x3b9aca00",
		  "gas": "0x5208",
		  "to": "0xrecipientAddress",
		  "value": "0xde0b6b3a7640000",
		  "input": "0x",
		  "v": "0x1c",
		  "r": "0xSignatureRValue",
		  "s": "0xSignatureSValue",
		  "hash": "0xTransactionHash",
		  "blockHash": "0xBlockHash",
		  "blockNumber": "0x1b4",
		  "transactionIndex": "0x0"
		}
	  ]
	}`),
			want: []TxnEIP2718{
				{
					Type: func() *string {
						value := "0x0"
						return &value
					}(),
					Nonce:            "0x1",
					GasPrice:         "0x3b9aca00",
					Gas:              "0x5208",
					To:               "0xrecipientAddress",
					Value:            "0xde0b6b3a7640000",
					Input:            "0x",
					V:                "0x1c",
					R:                "0xSignatureRValue",
					S:                "0xSignatureSValue",
					Hash:             "0xTransactionHash",
					BlockHash:        "0xBlockHash",
					BlockNumber:      "0x1b4",
					TransactionIndex: "0x0",
				},
			},
		},
		{
			input: []byte(`{
			"number": "1",
  "transactions": [
    {
      "type": "0x1",
      "chainId": "0x1",
      "nonce": "0x1",
      "gasPrice": "0x3b9aca00",
      "gas": "0x5208",
      "to": "0xrecipientAddress",
      "value": "0xde0b6b3a7640000",
      "input": "0x",
      "accessList": [
        {
          "address": "0xContractAddress",
          "storageKeys": ["0xStorageKey1"]
        }
      ],
      "v": "0x1c",
      "r": "0xSignatureRValue",
      "s": "0xSignatureSValue",
      "hash": "0xTransactionHash",
      "blockHash": "0xBlockHash",
      "blockNumber": "0x1b4",
      "transactionIndex": "0x0"
    }
  ]
}`),
			want: []TxnEIP2930{
				{
					Type: func() *string {
						value := "0x1"
						return &value
					}(),
					ChainID:  "0x1",
					Nonce:    "0x1",
					GasPrice: "0x3b9aca00",
					Gas:      "0x5208",
					To:       "0xrecipientAddress",
					Value:    "0xde0b6b3a7640000",
					Input:    "0x",
					AccessList: []AccessList{
						{
							Address:     "0xContractAddress",
							StorageKeys: []string{"0xStorageKey1"},
						},
					},
					V:                "0x1c",
					R:                "0xSignatureRValue",
					S:                "0xSignatureSValue",
					Hash:             "0xTransactionHash",
					BlockHash:        "0xBlockHash",
					BlockNumber:      "0x1b4",
					TransactionIndex: "0x0",
				},
			},
		},
		{
			input: []byte(`{
			"number": "1",
				"transactions": [
				  {
					"type": "0x2",
					"chainId": "0x1",
					"nonce": "0x1",
					"maxPriorityFeePerGas": "0x3b9aca00",
					"maxFeePerGas": "0x77359400",
					"gas": "0x5208",
					"to": "0xrecipientAddress",
					"value": "0xde0b6b3a7640000",
					"input": "0x",
					"v": "0x1c",
					"r": "0xSignatureRValue",
					"s": "0xSignatureSValue",
					"hash": "0xTransactionHash",
					"blockHash": "0xBlockHash",
					"blockNumber": "0x1b4",
					"transactionIndex": "0x0"
				  }
				]
			  }`),
			want: []TxnEIP1559{
				{
					Type: func() *string {
						value := "0x2"
						return &value
					}(),

					ChainID:              "0x1",
					Nonce:                "0x1",
					MaxPriorityFeePerGas: "0x3b9aca00",
					MaxFeePerGas:         "0x77359400",
					Gas:                  "0x5208",
					To:                   "0xrecipientAddress",
					Value:                "0xde0b6b3a7640000",
					Input:                "0x",
					V:                    "0x1c",
					R:                    "0xSignatureRValue",
					S:                    "0xSignatureSValue",
					Hash:                 "0xTransactionHash",
					BlockHash:            "0xBlockHash",
					BlockNumber:          "0x1b4",
					TransactionIndex:     "0x0",
				},
			},
		},
		{
			input: []byte(`{
  "transactions": [
    {
      "type": "0x5",
      "chainId": "0x1",
      "nonce": "0x1",
      "maxPriorityFeePerGas": "0x3b9aca00",
      "maxFeePerGas": "0x77359400",
      "gas": "0x5208",
      "to": "0xrecipientAddress",
      "value": "0xde0b6b3a7640000",
      "input": "0x",
      "blobVersionedHashes": [
        "0xBlobHash1",
        "0xBlobHash2"
      ],
      "v": "0x1c",
      "r": "0xSignatureRValue",
      "s": "0xSignatureSValue",
      "hash": "0xTransactionHash",
      "blockHash": "0xBlockHash",
      "blockNumber": "0x1b4",
      "transactionIndex": "0x0"
    }
  ]
}`),
			want: []TxnEIP4844{
				{
					Type: func() *string {
						value := "0x5"
						return &value
					}(),
					ChainID:              "0x1",
					Nonce:                "0x1",
					MaxPriorityFeePerGas: "0x3b9aca00",
					MaxFeePerGas:         "0x77359400",
					Gas:                  "0x5208",
					To:                   "0xrecipientAddress",
					Value:                "0xde0b6b3a7640000",
					Input:                "0x",
					BlobVersonedHashes:   []string{"0xBlobHash1", "0xBlobHash2"},
					V:                    "0x1c",
					R:                    "0xSignatureRValue",
					S:                    "0xSignatureSValue",
					Hash:                 "0xTransactionHash",
					BlockHash:            "0xBlockHash",
					BlockNumber:          "0x1b4",
					TransactionIndex:     "0x0",
				},
			},
		},
	}

	for i, tc := range testcases {
		var blk Block
		err := json.Unmarshal(tc.input, &blk)
		if assert.NoError(t, err, fmt.Sprintf("Case: %d Error: %v", i, err)) {
			switch got := blk.Transactions.(type) {
			case []string:
				assert.Equal(t, tc.want, got, fmt.Sprintf("Case: %d Want: %v Got: %v", i, tc.want, got))
			case []TxnPreEIP2718:
				assert.Equal(t, tc.want, got, fmt.Sprintf("Case: %d Want: %v Got: %v", i, tc.want, got))
			case []TxnEIP2718:
				assert.Equal(t, tc.want, got, fmt.Sprintf("Case: %d Want: %v Got: %v", i, tc.want, got))
			case []TxnEIP2930:
				assert.Equal(t, tc.want, got, fmt.Sprintf("Case: %d Want: %v Got: %v", i, tc.want, got))
			case []TxnEIP1559:
				assert.Equal(t, tc.want, got, fmt.Sprintf("Case: %d Want: %v Got: %v", i, tc.want, got))
			case []TxnEIP4844:
				assert.Equal(t, tc.want, got, fmt.Sprintf("Case: %d Want: %v Got: %v", i, tc.want, got))
			default:
				assert.Fail(t, "Case not caught", fmt.Sprintf("Case: %d", i))
			}
		}
	}
}

func TestTransformTxnArg(t *testing.T) {
	testcases := []struct {
		input TxnArg
		want  map[string]any
	}{
		{
			input: TxnArg{
				To:   "Hello",
				From: "World",
			},
			want: map[string]any{"from": "World", "to": "Hello"},
		},
		{
			input: TxnArg{
				AccessList: []AccessListArg{
					{
						Address:     "<address>",
						StorageKeys: []string{"abc", "efg"},
					},
				},
			},
			want: map[string]any{"accessList": []any{map[string]any{"address": "<address>", "storageKeys": []any{"abc", "efg"}}}},
		},
	}

	for i, tc := range testcases {
		got, err := transformTxnArg(tc.input)
		if assert.NoError(t, err, fmt.Sprintf("Case: %d Error: %v", i, err)) {
			assert.Equal(t, tc.want, got, fmt.Sprintf("Case: %d Want: %v Got: %v", i, tc.want, got))
		}
	}
}
