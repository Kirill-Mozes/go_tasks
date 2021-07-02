package main

import (
	"fmt"
	"time"
)

//Написать собственную функцию Sleep

func Sleep(seconds int) {
	if seconds > 0 {
		timer := time.NewTimer(time.Duration(seconds) * time.Second)
		<-timer.C
	}
}
func prin() {
	Sleep(2)
	fmt.Println("bib")
}
func main() {
	fmt.Println("start")
	go prin()
	Sleep(15)

	fmt.Println("end")
}
