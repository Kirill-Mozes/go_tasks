package main

//Написать пересечение двух неупорядоченных массивов
import (
	"fmt"
)

func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1 // если встретили в первый раз
		} else {
			counter[elem] += 1
		}
	}
	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 { //если есть в мапе пишем в финальный массив
			counter[elem] -= 1
			result = append(result, elem)
		}
	}
	return result
}
func main() {
	a := []int{23, 3, 1, 2, 2}
	b := []int{6, 2, 4, 23, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}
