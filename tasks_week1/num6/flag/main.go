package main

import (
	"fmt"
	"time"
)

func f(flag *bool) {
	for {
		if *flag {
			fmt.Println(1)
			return
		}
	}

}
func main() {

	flag := false
	go f(&flag)

	time.Sleep(time.Second)
	flag = true
	time.Sleep(time.Second)
	fmt.Println(2)

}
