package main

import "fmt"

func main() {
	x := 1
	y := 6
	x, y = y, x // go style
	fmt.Println("x: ", x, " y:", y)

	a := 10
	b := 4
	a = a + b //c ctyle
	b = a - b
	a = a - b
	fmt.Println("a: ", a, "b: ", b)
}
