package main

import (
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/wallet"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Create Mnemonic
	m, err := wallet.Mnemonic(wallet.Mnemonic128)
	if err != nil {
		log.Fatal(err)
	}

	// Create master key with a combination of Mnemonic and passphrase
	passphrase := "hello"
	mkey, err := wallet.MasterHDKey(m, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type: %[1]T Value: %[1]v\n", mkey.Key)

	// Generate first child key
	childKey, err := mkey.NewChildKey(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(childKey)

	// Recover key from Mnemonic and passphrase
	mkey1, err := wallet.MasterHDKey(m, passphrase)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %[1]T Value: %[1]v\n", mkey1.Key)

	mkeyHex := hexutil.Encode(mkey.Key)
	mkey1Hex := hexutil.Encode(mkey1.Key)
	fmt.Println(mkeyHex, mkey1Hex)

	// Convert private key in bytes to ECDSA
	privkey, err := crypto.ToECDSA(mkey.Key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T", privkey)

}
