package contract

import (
	"context"
	"crypto/ecdsa"
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

func extractContractBin(binFile string) (string, error) {
	data, err := os.ReadFile(binFile)
	if err != nil {
		return "", err
	}
	content := fmt.Sprintf("0x%v", string(data))
	return content, nil
}

// DeployContract is an operation to deploy contract
func DeployContract(ctx context.Context, client *ethclient.Client, privKey *ecdsa.PrivateKey, gasLimit uint64, compiledContract string) (common.Address, error) {
	return deployContract(ctx, client, privKey, gasLimit, compiledContract)
}

func deployContract(ctx context.Context, client *ethclient.Client, privKey *ecdsa.PrivateKey, gasLimit uint64, compiledContract string) (common.Address, error) {
	publicKey := privKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return common.Address{}, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Address{}, err
	}

	// Smart contract bytecode and ABI (compiled using solc or Remix IDE)
	tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, common.FromHex(compiledContract))

	// Retrieve the ID of the network to which the client is connected
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return common.Address{}, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return common.Address{}, err
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Address{}, err
	}

	// Get contract address
	contractAddress := crypto.CreateAddress(fromAddress, nonce)

	return contractAddress, nil
}
