package main

//Написать объединение двух строк
import (
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func unification(a, b string) string {
	counter := make(map[rune]int)
	var result []rune

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1 // если встретили в первый раз
		} else {
			counter[elem] += 1
		}
	}
	for _, elem := range b {
		if _, ok := counter[elem]; !ok {
			counter[elem] = -1 // если встретили в первый раз
		} else if _, ok := counter[elem]; ok { //если есть в мапе пишем в финальный массив
			counter[elem] -= 1
			//result = append(result, elem)
		}
	}
	for key, value := range counter {
		for i := 0; i < abs(value); i++ {
			result = append(result, key)
		}
	}
	return string(result)
}
func main() {
	a := "aasaab"
	b := "aaaabb"
	fmt.Printf("%v\n", unification(a, b))
}
