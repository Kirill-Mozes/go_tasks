package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1 // это другая n первоначальное присваивание области видимости
		n++
	}
	fmt.Println(n)
} //0
