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
