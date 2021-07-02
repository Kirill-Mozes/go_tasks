package main

import (
	"fmt"
	"math/big"
)

func sum(a, b *big.Int) *big.Int {
	c := new(big.Int)
	return c.Add(a, b)
}
func multiply(a, b *big.Int) *big.Int {
	c := new(big.Int)
	return c.Mul(a, b)
}
func minus(a, b *big.Int) *big.Int {
	c := new(big.Int)
	return c.Sub(a, b)
}
func divide(a, b *big.Int) *big.Int {
	c := new(big.Int)
	return c.Div(a, b)
}
func main() {

	a := big.NewInt(100000000)
	b := big.NewInt(50000)
	c := new(big.Int)
	c = sum(a, b)
	fmt.Println(c)
	fmt.Println(sum(a, b))
	fmt.Println(multiply(a, b))
	fmt.Println(minus(a, b))
	fmt.Println(divide(a, b))
}
