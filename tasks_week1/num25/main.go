package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex // мьютекс
	i          int
}

func New() *Counter {
	return &Counter{
		i: 0,
	}
}

func (c *Counter) Count() {
	c.Lock()
	defer c.Unlock()
	c.i += 1
	fmt.Println(c.i)
	//return c.i
}

func main() {
	store := New()
	for i := 0; i < 15; i++ {
		go func() {
			store.Count()
			//fmt.Println(counter.i)
		}()
	}
	fmt.Scanln()
}
