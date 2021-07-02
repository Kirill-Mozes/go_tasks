package main

import "fmt"

func isUnique(str string) bool {
	m := make(map[rune]int)
	for _, v := range str {
		m[v] += 1
	}
	for _, value := range m {
		if value > 1 {
			return false
		}
	}
	return true
}
func main() {
	a := ""
	fmt.Println(isUnique((a)))
}
