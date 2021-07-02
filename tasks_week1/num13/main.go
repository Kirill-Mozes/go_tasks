package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) { //wg *sync.WaitGroup нужен указатель
			fmt.Println(i)
			wg.Done()
		}(&wg, i) //&wg и тут
	}
	wg.Wait()
	fmt.Println("exit")
}

/*4
2
1
3
0
fatal error: all goroutines are asleep - deadlock!
*/
