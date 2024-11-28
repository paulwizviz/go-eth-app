package keystore

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func Create(ksdir string, passphrase string) (accounts.Account, error) {
	return createKeystore(ksdir, passphrase)
}

func createKeystore(ksdir string, passphrase string) (accounts.Account, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return accounts.Account{}, err
	}
	ks := keystore.NewKeyStore(ksdir, keystore.StandardScryptN, keystore.StandardScryptP)
	acct, err := ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Keystore account:", acct.Address.Hex())

	return accounts.Account{}, nil
}
