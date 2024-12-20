package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/paulwizviz/go-eth-app/internal/contract/hello"
)

func main() {
	// Generate a private key for testing
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // Chain ID for testing
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Initialize SimulatedBackend with prefunded accounts
	genesisAlloc := core.GenesisAlloc{
		auth.From: {Balance: big.NewInt(1e18)}, // 1 Ether
	}
	client := backends.NewSimulatedBackend(genesisAlloc, 8000000) // Gas limit

	// Deploy the contract
	initialValue := big.NewInt(42)
	address, _, instance, err := hello.DeployHelloWorld(auth, client, initialValue)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	client.Commit() // Commit the transaction to the simulated blockchain
	log.Printf("Contract deployed at address: %s\n", address.Hex())

	// Interact with the contract
	value, err := instance.GetValue(nil)
	if err != nil {
		log.Fatalf("Failed to call getValue: %v", err)
	}
	log.Printf("Initial Stored Value: %d\n", value)

	// Mutate state using setValue
	auth.Nonce = big.NewInt(1) // Increment nonce for the next transaction
	tx, err := instance.SetValue(auth, big.NewInt(100))
	if err != nil {
		log.Fatalf("Failed to call setValue: %v", err)
	}
	client.Commit() // Commit the transaction

	fmt.Println(tx)

	// Verify the state change
	value, err = instance.GetValue(nil)
	if err != nil {
		log.Fatalf("Failed to call getValue: %v", err)
	}
	log.Printf("Updated Stored Value: %d\n", value)
}
