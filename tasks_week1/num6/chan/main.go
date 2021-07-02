package main

//Какие существуют способы остановить выполнения горутины
import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok { //пока канал не закрыт
				println("done")
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	wg.Wait()
	/*
		quit := make(chan bool)
		go func() {
			for {
				select {
				case <-quit:
					return
				default:
					// Do other stuff
				}
			}
		}()

		// Do stuff

		// Quit goroutine
		quit <- true*/
}
