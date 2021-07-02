package main

import (
	"fmt"
	_ "sync"
)

//Написать программу, которая в конкурентном виде читает элементы из массива в stdout
func main() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	//var mux sync.Mutex
	for _, v := range arr {
		go func(v int) {
			//mux.Lock()
			//defer mux.Unlock()
			fmt.Println(v)
		}(v)
	}
	fmt.Scanln()
}
