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

package ether

import (
	"math/big"
)

const (
	kwei          = int64(1_000)                     // wei
	mwei          = int64(1_000_000)                 // wei
	gwei          = int64(1_000_000_000)             // wei
	microetherWei = int64(1_000_000_000_000)         // wei
	millietherWei = int64(1_000_000_000_000_000)     // wei
	etherWei      = int64(1_000_000_000_000_000_000) // wei

	etherGwei = int64(1_000_000_000) // gwei
)

// Gas is a unit of computation
type Gas uint64

// Wei is the smallest unit of Ether
type Wei int64

// ToGweiBF returns Wei in big.Float
func (w Wei) ToGweiBF() *big.Float {
	numerator := big.NewFloat(float64(w))
	denominator := big.NewFloat(float64(gwei))
	result := new(big.Float)
	return result.Quo(numerator, denominator)
}

// ToEtherBF returns Ether in big.Float
func (w Wei) ToEtherBF() *big.Float {
	convFactor := big.NewFloat(float64(etherWei))
	numerator := big.NewFloat(float64(w))
	result := new(big.Float)
	return result.Quo(numerator, convFactor)
}

// Gwei is 1,000,000,000 wei
type Gwei float64

// ToWeiBI returns Wei in Big.Int and accuracy
func (g Gwei) ToWeiBI() (*big.Int, big.Accuracy) {
	conversionFactor := big.NewFloat(float64(gwei))
	multiplicant := big.NewFloat(float64(g))
	result := new(big.Float)
	return result.Mul(multiplicant, conversionFactor).Int(nil)
}

// ToEther returns Ether big.Float
func (g Gwei) ToEther() *big.Float {
	conversionFactor := big.NewFloat(float64(etherGwei))
	nominator := big.NewFloat(float64(g))
	result := new(big.Float)
	return result.Quo(nominator, conversionFactor)
}

// Ether is 1,000,000,000,000,000,000 wei
// or 1,000,000,000 gwei
type Ether float64

// ToWeiBI returns wei in big.Int and accuracy in big.Accuracy
func (e Ether) ToWeiBI() (*big.Int, big.Accuracy) {
	conversionFactor := big.NewFloat(float64(etherWei))
	multiplicant := big.NewFloat(float64(e))
	return new(big.Float).Mul(multiplicant, conversionFactor).Int(nil)
}

// ToGweiBI returns gwei in big.Int and accuracy in big.Accuracy
func (e Ether) ToGweiBI() (*big.Int, big.Accuracy) {
	conversionFactor := big.NewFloat(float64(etherGwei))
	multiplicant := big.NewFloat(float64(e))
	return new(big.Float).Mul(multiplicant, conversionFactor).Int(nil)
}
