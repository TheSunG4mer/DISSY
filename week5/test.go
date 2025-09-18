package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(3), big.NewInt(2)
	c := big.NewInt(0)
	c.Mul(a, b)
	fmt.Println(c)
}
