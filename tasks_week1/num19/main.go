package main

import (
	"fmt"
	"unsafe"
)

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "A"
	}
	return v
}

var justString string
var otherString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:99]                   //это не 99 символов а 99 байт поэтому строка бьется
	otherString := string([]rune(v)[:99]) //исправление
	//c := make([]rune, 99)
	//c += copy(c[0:], "f")
	//fmt.Println("cap  v: ", cap([]byte(v)), " len  v: ", len([]byte(v)), "size v: ", unsafe.Sizeof(v))
	//fmt.Println("cap  j: ", cap([]byte(justString)), " len  j: ", len([]byte(justString)), "size j: ", unsafe.Sizeof(justString))
	//fmt.Println("cap  o: ", cap([]byte(otherString)), " len  o: ", len([]byte(otherString)), "size o: ", unsafe.Sizeof(otherString))
	fmt.Println(&v)
	fmt.Println(&justString)
	fmt.Println(&otherString)
	fmt.Println(&c)
}

func main() {
	someFunc()
	fmt.Println("cap  j: ", cap([]byte(justString)), " len  j: ", len([]byte(justString)), "size j: ", unsafe.Sizeof(justString))
}
