package main

import (
	"math/big"
)

func factorial(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}
	previousFactorial := factorial(n - 1)
	current := big.NewInt(int64(n))
	return current.Mul(current, previousFactorial)
}
