package main

import "fmt"

//Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов с использованием конкурентных вычислений
func square(value int, c chan int) {
	c <- value * value //wiriting in chan
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10} // массив
	c := make(chan int)
	for _, v := range arr {
		go square(v, c) // конкурентные вычисления
	}
	count := 0
	for i := 0; i < len(arr); i++ {
		a := <-c
		count += a
	}
	fmt.Println(count) //pinting in stdout
	fmt.Scanln()

}
