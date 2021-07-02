package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100 //меняет внешнюю
	fmt.Println(cap(v), &v[0])
	v = append(v, b) //не меняет внешнюю, тк меняется ссылка на участок памяти тк изменился размер
	v[2] = 100
	fmt.Println(cap(v), &v[0]) //чтд
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a, &a[0])
	//fmt.Println(len(a))
}

//[100 2 3 4 5 ]
