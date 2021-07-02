package main

//Даны 2 канала - в первый пишутся рандомные числа после чего они проверяются на четность и отправляются во второй канал.
//Результаты работы из второго канала пишутся в stdout
import (
	"fmt"
	"math/rand"
)

func main() {

	randoms := make(chan int)
	multipying := make(chan int)

	go func() {
		for x := 0; x <= 10; x++ {
			num := rand.Intn(33)
			randoms <- num
		}
		close(randoms)
	}()

	go func() {
		for x := range randoms {
			if x%2 == 0 {
				multipying <- x
			}
		}
		close(multipying)
	}()

	for x := range multipying {
		fmt.Println(x)
	}
}
