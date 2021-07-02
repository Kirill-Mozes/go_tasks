package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { //сходимся к центру
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	str := "ф24ф日本п語gjQ"
	fmt.Println(str)
	fmt.Println(Reverse(str))
}
