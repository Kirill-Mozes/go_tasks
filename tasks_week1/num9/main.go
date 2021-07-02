package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	multipying := make(chan int)
	arr := [8]int{1, 2, 3, 4, 6, 8, 10, 5}
	go func(ar []int) { //срез для универсальности
		for x := 0; x <= len(ar)-1; x++ {
			naturals <- ar[x] //пишем в канал
		}
		close(naturals)
	}(arr[:])

	go func() {
		for x := range naturals {
			multipying <- 2 * x
		}
		close(multipying)
	}()

	for x := range multipying {
		fmt.Println(x)
	}
}
