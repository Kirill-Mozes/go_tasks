package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println(str)
	words := strings.Fields(str) //делим по пробелам
	str = ""
	for i := len(words) - 1; i >= 0; i-- {
		str += (words[i] + " ")
	}
	fmt.Println(str)
}
