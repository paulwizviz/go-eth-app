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
