package main

import (
	"fmt"
	"sync"
	_ "sync"
)

func main() {
	var counters = map[int]int{}
	var mux sync.Mutex
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int, mux *sync.Mutex) { // обращаемся к мапе из разных потоков
			mux.Lock()
			defer mux.Unlock()
			for j := 0; j < 5; j++ {
				counters[th*10+j]++
			}
		}(counters, i, &mux) //
	}
	fmt.Scanln()
	fmt.Println("counters result ", counters)
}
