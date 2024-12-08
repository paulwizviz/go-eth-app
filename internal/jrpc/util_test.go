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

package jrpc

import (
	"math/big"
	"testing"
)

var (
	want = "0x30a8"
)

func TestIntToStringHex(t *testing.T) {
	got := IntToHexString(int(12456))
	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}

func TestBigIntToStringHex(t *testing.T) {
	input := big.NewInt(12456)
	got := IntToHexString(input)
	if got != want {
		t.Errorf("Want: %s Got: %s", want, got)
	}
}
