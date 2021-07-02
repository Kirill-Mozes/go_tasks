package main

import "fmt"

func update(p *int) {
	b := 2
	fmt.Println("b in func1= ", &b)
	fmt.Println("p in func1= ", p)
	p = &b //это другое p
	fmt.Println("p in func2= ", p)
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p) //print a
	fmt.Println("p in main1= ", p)
	update(p)
	fmt.Println("p in main2= ", p)
	fmt.Println(*p)
}

//1
//1
