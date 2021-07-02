package main

import (
	"os"
	"time"
)

func foo() {
	i := 0
	for {
		i++
	}
}

func main() {
	go foo()
	time.Sleep(3 * time.Second)
	os.Exit(1)
}
