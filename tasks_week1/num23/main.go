package main

import "fmt"

func binarySearch(arr []int, value int) (string, int) {
	flag := false
	l := 0
	r := len(arr) - 1
	var mid int
	for (l <= r) && (flag != true) {
		mid = (l + r) / 2
		if arr[mid] == value {
			flag = true
		} else if arr[mid] > value {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	if flag {
		return "your element ", value
	} else {
		return "not find", 0
	}
}

func main() {
	values := []int{1, 2, 3, 4, 6, 7}
	fmt.Println(binarySearch(values, 5))
}
