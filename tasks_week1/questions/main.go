package main

import "fmt"

//Написать объединение двух строк

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func unification(a, b string) string {
	counterA := make(map[rune]int)
	counterB := make(map[rune]int)
	var result []rune

	for _, elem := range a { //алфавит строки а
		if _, ok := counterA[elem]; !ok {
			counterA[elem] = 1 // если встретили в первый раз
		} else {
			counterA[elem] += 1
		}
	}

	for _, elem := range b { //алфавит строки б
		if _, ok := counterB[elem]; !ok {
			counterB[elem] = 1 // если встретили в первый раз
		} else {
			counterB[elem] += 1
		}
	}
	for key, value := range counterA { //объединяем алфавиты
		if _, ok := counterB[key]; !ok {
			counterB[key] = value // если встретили в первый раз
		} else {
			counterB[key] = abs(counterB[key] - counterA[key]) //пересечение
		}
	}
	for key, value := range counterB {
		for i := 0; i < value; i++ {
			result = append(result, key)
		}
	}
	return string(result)
}
func main() {
	a := "aaab"
	b := "aaaab"
	fmt.Printf("%v\n", unification(a, b))
}
