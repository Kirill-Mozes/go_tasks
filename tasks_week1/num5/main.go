package main

import (
	"fmt"
	"os"
	"time"
)

const n = 5

func programmStop() {
	time.Sleep(n * time.Second)
	os.Exit(1) //останавливаем программу
}

func writing(c chan int) {
	i := 0
	for true {
		c <- i
		i++
	}

	close(c) // close channel
}

func main() {
	go programmStop()
	c := make(chan int)

	go writing(c) // start goroutine

	for {
		val, ok := <-c
		if ok == false {
			fmt.Println("broke")
			break // exit break loop
		} else {
			fmt.Println(val)
		}
	}

	fmt.Println("stop")
}

/*
package main
*/
//Написать программу, которая будет последовательно писать значения в канал,
//а с другой стороны канала — читать.
//По истечению N секунд программа должна завершиться.
/*
import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const n = 5

func write(tasksCh chan int, i int) {
	tasksCh <- i
}
func read(tasksCh <-chan int) {
	value := <-tasksCh
	fmt.Println(value)
}
func programmStop() {
	time.Sleep(n * time.Second)
	os.Exit(1)
}
func main() {
	go programmStop()
	var values int
	valuesCh := make(chan int, values)
	defer close(valuesCh)

	for true {
		values := rand.Intn(1000)
		go write(valuesCh, values)
		go read(valuesCh)
	}

}
*/
