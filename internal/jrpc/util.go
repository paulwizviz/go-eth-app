package jrpc

import (
	"fmt"
	"math/big"
)

func BigIntToHexString(num *big.Int) string {
	return fmt.Sprintf("0x%x", num)
}

func IntToHexString[T int | *big.Int](num T) string {
	return fmt.Sprintf("0x%x", num)
}
