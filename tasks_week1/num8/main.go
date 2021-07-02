package main

import (
	"fmt"
	"os"
)

//Дана переменная int64. Написать программу которая устанавливает i-й бит в 1 или 0

func main() {

	var x int64
	x = 1
	var b int
	var number int64
	number = 2
	fmt.Println("what bit")
	fmt.Scanf("%d\n", &b)
	if b > 62 {
		//закончить прогамму
		fmt.Println("not valid number")
		os.Exit(1)
	}
	var want uint8
	fmt.Println("0 or 1")
	fmt.Scanf("%d\n", &want)
	fmt.Println(x)
	if want == 1 {
		x = x << b
		number = number | x //установить в еденицу
		fmt.Println("number:", number)
	} else if want == 0 {
		x = x << b
		number = number &^ x //установить в ноль
		fmt.Println("number:", number)
	} else {
		os.Exit(1)
	}

}
