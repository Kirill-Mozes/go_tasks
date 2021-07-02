package main

import "fmt"

func binarySearch(array []int, searchedValue, first, last int) int {
	//границы сошлись
	if first > last {
		//элемент не найден
		return -1
	}

	//средний индекс подмассива
	middle := (first + last) / 2
	//значение в средине подмассива
	middleValue := array[middle]

	if middleValue == searchedValue {
		return array[middle]
	} else {
		if middleValue > searchedValue {
			//рекурсивный вызов поиска для левого подмассива
			return binarySearch(array, searchedValue, first, middle-1)
		} else {
			//рекурсивный вызов поиска для правого подмассива
			return binarySearch(array, searchedValue, middle+1, last)
		}
	}
}
func main() {
	values := []int{1, 2, 3, 4, 6, 7}
	fmt.Println("your element: ", binarySearch(values, 3, 0, len(values)-1))
}
