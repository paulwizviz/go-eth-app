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
	"errors"
	"fmt"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

var (
	ErrMnemonic = errors.New("Mnemonic error")
)

const (
	Mnemonic128 = 128
	Mnemonic256 = 256
)

// Mnemonic takes a secret phrase and bitsize argemnents to generate
// a mnemonic if no error ecountered.
func Mnemonic(bitsize int) (string, error) {
	entropy, err := bip39.NewEntropy(bitsize)
	if err != nil {
		return "", fmt.Errorf("%w-%s", ErrMnemonic, err.Error())
	}
	m, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("%w-%s", ErrMnemonic, err.Error())
	}
	return m, nil
}

// MnemonicEntropy generates an entropy
func MnemonicEntropy(entropy string) (string, error) {
	m, err := bip39.NewMnemonic([]byte(entropy))
	if err != nil {
		return "", nil
	}
	return m, nil
}

// MasterHDKey takes a mnemonic and passphrase arguments to return a bip32 key
// if no error encountered.
func MasterHDKey(mnemonic string, passphrase string) (*bip32.Key, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	mkey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}
	return mkey, nil
}
