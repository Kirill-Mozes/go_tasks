package main

//Написать программу, которая конкурентно рассчитает значение квадратов
//значений взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
)

func square(value int) {

	fmt.Println(value * value) //pinting in stdout

}

func main() {
	arr := [5]int{2, 4, 6, 8, 10} //array
	for _, v := range arr {
		go square(v) //конкурентныйы рассчет
	}
	fmt.Scanln()
}
