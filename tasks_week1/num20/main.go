package main

import "fmt"

/*
func main() {
	slice := []string{"a", "a"}
	fmt.Println("str7 ", cap(slice), &slice[0])
	func(slice []string) {
		fmt.Println("str9 ", cap(slice), &slice[0])
		slice = append(slice, "a")
		fmt.Println("str11 ", cap(slice), &slice[0])
		slice[0] = "b"
		slice[1] = "b"
	}(slice)
	fmt.Print(cap(slice), &slice[0])
}*/
func main() {
	slice := []string{"a", "a"}

	func(slice []string) { //передаем по значению
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}

//аналогично что и в номере 18 изменяется участок памяти тк копируется массив в новую область
//[b b a][a a]
