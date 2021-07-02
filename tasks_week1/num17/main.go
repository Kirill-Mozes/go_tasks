package main

import (
	"fmt"
	"reflect"
)

// Написать программу, которая в рантайме способна определить тип переменной
// int, string, bool, channel из переменной типа interface{}.

func detector(v interface{}) {
	//fmt.Println("type:", reflect.TypeOf(v))
	a := reflect.TypeOf(v).String()

	switch a {
	case "int":
		fmt.Println("int")
	case "bool":
		fmt.Println("bool")
	case "string":
		fmt.Println("string")
	case "chan int":
		fmt.Println("chanal")
	}
}
func main() {
	detector("dff")
	detector(323)
	detector(true)
	detector(make(chan int))
	a := 5
	detector(a)
}
