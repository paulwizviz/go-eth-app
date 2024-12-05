package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ExtractContent exteact the content of bin file
func ExtractContent(binFile string) (string, error) {
	return extractContractBin(binFile)
}

// ErrUnableToExtractContent error extracting contract content
var ErrUnableToExtractContent = errors.New("unable to extract contract error")

func extractContractBin(binFile string) (string, error) {
	data, err := os.ReadFile(binFile)
	if err != nil {
		return "", fmt.Errorf("%w-%v", ErrUnableToExtractContent, err)
	}
	content := fmt.Sprintf("0x%v", string(data))
	return content, nil
}

// DeployContract is an operation to deploy contract
func DeployContract(ctx context.Context,
	client *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	gasTipCap int64,
	gasFeeCap int64,
	gasLimit uint64,
	compiledContract string) (common.Address, error) {
	return deployContract(ctx, client, privKey, gasTipCap, gasFeeCap, gasLimit, compiledContract)
}

var (
	// ErrUnableToGetPendingNonce error trying to get pending nonce
	ErrUnableToGetPendingNonce = errors.New("unable to get pending nonce error")
	// ErrUnableToGetSuggestedGasPrice error trying to get suggested gas price
	ErrUnableToGetSuggestedGasPrice = errors.New("unable to get suggested gas price error")
	// ErrUnableToConvertContractToHex error converting contract content to Hex
	ErrUnableToConvertContractToHex = errors.New("unable to convert contract content to Hex error")
	// ErrUnableToGetNetworkID error getting network ID
	ErrUnableToGetNetworkID = errors.New("unable to get network ID error")
	// ErrUnableToSignTxn
	ErrUnableToSignTxn = errors.New("unable to sign transaction error")
	// ErrUnableToSendTxn error sending transaction
	ErrUnableToSendTxn = errors.New("unable to send txn error")
)

func deployContract(ctx context.Context,
	client *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	gasTipCap int64,
	gasFeeCap int64,
	gasLimit uint64,
	compiledContract string) (common.Address, error) {

	publicKey := privKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("%w-%v", ErrUnableToGetPendingNonce, err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Address{}, fmt.Errorf("%w-%v", ErrUnableToGetSuggestedGasPrice, err)
	}

	contractBytes, err := hex.DecodeString(compiledContract)
	if err != nil {
		return common.Address{}, fmt.Errorf("%w-%v", ErrUnableToConvertContractToHex, err)
	}

	// Retrieve the ID of the network to which the client is connected
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return common.Address{}, fmt.Errorf("%w-%v", ErrUnableToGetNetworkID, err)
	}

	// Smart contract bytecode and ABI (compiled using solc or Remix IDE)
	txData := &types.DynamicFeeTx{
		ChainID:   chainID, // Chain ID for the network
		Nonce:     nonce,
		GasTipCap: big.NewInt(gasTipCap),                             // Tip
		GasFeeCap: new(big.Int).Add(gasPrice, big.NewInt(gasFeeCap)), // Base fee + tip
		Gas:       gasLimit,
		To:        nil,           // `To` is nil for contract deployment
		Value:     big.NewInt(0), // Value sent with the transaction
		Data:      contractBytes, // Contract bytecode
	}

	tx := types.NewTx(txData)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return common.Address{}, fmt.Errorf("%w-%v", ErrUnableToSignTxn, err)
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Address{}, fmt.Errorf("%w-%v", ErrUnableToSendTxn, err)
	}

	// Get contract address
	contractAddress := crypto.CreateAddress(fromAddress, nonce)

	return contractAddress, nil
}
