package main

import "fmt"

//Создать слайс с предварительно выделенными 100 элементами
func main() {
	n := make([]int, 0, 100)
	fmt.Println("cap: ", cap(n))
	fmt.Println("len: ", len(n))
	fmt.Println("arr: ", n)
}
