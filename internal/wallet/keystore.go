package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// KeystoreCreate instantiate an instance of a key store wallet in path specified.
func KeystoreCreate(kspath string, passphrase string) (accounts.Account, error) {
	return keystoreCreate(kspath, passphrase)
}

func keystoreCreate(kspath string, passphrase string) (accounts.Account, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return accounts.Account{}, fmt.Errorf("%w-%v", ErrKeystoreGenerateKey, err)
	}
	ks := keystore.NewKeyStore(kspath, keystore.StandardScryptN, keystore.StandardScryptP)
	acct, err := ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("%w-%v", ErrKeystoreCreate, err)
	}
	return acct, nil
}

// KeystoreRecoverPrivKey recover private key from keystore file
func KeystoreRecoverPrivKey(ksfile string, passphrase string) (*ecdsa.PrivateKey, error) {
	return keystoreRecoverPrivKey(ksfile, passphrase)
}

func keystoreRecoverPrivKey(kspath string, passphrase string) (*ecdsa.PrivateKey, error) {
	keystoreJSON, err := os.ReadFile(kspath)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrKeystoreReadFile, err)
	}

	key, err := keystore.DecryptKey(keystoreJSON, passphrase)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrKeystoreDecrypt, err)
	}
	return key.PrivateKey, nil
}
