package main

import (
	"fmt"
)

//Написать быструю сортировку встроенными методами языка.

func partition(a []int, l, r int) int {
	pivot := a[r]
	for j := l; j < r; j++ {
		if a[j] < pivot {
			a[j], a[l] = a[l], a[j]
			l++
		}
	}

	a[l], a[r] = a[r], a[l]
	return l
}

func quickSort(a []int, l, r int) {
	if l > r {
		return
	}

	p := partition(a, l, r)
	quickSort(a, l, p-1)
	quickSort(a, p+1, r)
}
func main() {
	//list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}
	list := []int{55, 53, 41, 93, 95, 84, 88}
	fmt.Println(list)
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)

}
