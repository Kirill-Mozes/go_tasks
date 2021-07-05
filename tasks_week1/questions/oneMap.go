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
	counter := make(map[rune]bool)
	var result []rune
	for _, elem := range a {
		result = append(result, elem) //сразу заполняем массивом а
		if _, ok := counter[elem]; !ok {
			counter[elem] = true //создаем мапу элементов
		}
	}
	for _, elem := range b {
		if _, ok := counter[elem]; !ok {
			// если такого элемента не было в прошлом массиве
			result = append(result, elem)
		}
	}

	return string(result)
}
func main() {
	a := "ggaasaab"
	b := "aasaabff"
	fmt.Printf("%v\n", unification(a, b)) //ggaasaabff
}
