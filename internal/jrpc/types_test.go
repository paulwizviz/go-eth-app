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
		assert.NoError(t, err, fmt.Sprintf("Case: %d Error: %v", i, err))

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
