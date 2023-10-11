package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("99000000000000000000", 10)

	b := big.NewInt(9900000000000)

	sum := new(big.Int)
	sum.Add(a, b)

	sub := new(big.Int)
	sub.Sub(a, b)

	mult := new(big.Int)
	mult.Mul(a, b)

	div := new(big.Int)
	div.Div(a, b)

	fmt.Printf("a=%v\nb=%v\nsum=%v\nsub=%v\nmult=%v\ndiv=%v\n", a, b, sum, sub, mult, div)
}
