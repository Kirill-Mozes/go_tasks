package main

//Реализовать набор из N воркеров, которые читают из канала произвольные данные и выводят в stdout.
//Данные в канал пишутся из главного потока.
//Необходима возможность выбора кол-во воркеров при старте, а также способ завершения работы всех воркеров.

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup, mtx *sync.Mutex, stop chan bool) {
	defer wg.Done()
	mtx.Lock()
	defer mtx.Unlock()
	for {
		select {
		case <-stop:
			fmt.Println("stop")
			return

		default:
			task, ok := <-tasksCh
			if !ok {
				return
			}
			fmt.Println("processing task", task)
		}
	}

}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int, tasks) //create channel
	stop := make(chan bool)
	var mux sync.Mutex
	//go func(stop chan bool) {
	//	stop <- true
	//}(stop)
	for i := 0; i < workers; i++ {
		time.Sleep(2 * time.Second)
		go worker(tasksCh, wg, &mux, stop) //workers
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i //write in channel
	}

	close(tasksCh)
}
func main() {
	fmt.Println("How many workwers do u want?")
	var i int
	fmt.Scanf("%d\n", &i)
	var wg sync.WaitGroup
	wg.Add(i) // определеяем wait group
	pool(&wg, i, 50)
	wg.Wait() // ждем завершения

}
