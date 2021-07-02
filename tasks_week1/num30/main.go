package main

import "fmt"

//Удалить i-ый элемент из слайса
func main() {

	//без сохранения порядка
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Удалить элемент по индексу i из a.
	fmt.Println(a)
	a[i] = a[len(a)-1] // Копировать последний элемент в индекс i
	a[len(a)-1] = ""   //Удалить последний элемент (записать нулевое значение)
	a = a[:len(a)-1]   // Усечь срез
	fmt.Println(a)     // [A B E D]

	//с сохранением порядка
	b := []string{"A", "B", "C", "D", "E"}
	j := 2

	// Удалить элемент по индексу i из a.
	copy(b[j:], b[j+1:]) // выполнить сдвиг a[i+1:] влево на один индекс
	b[len(b)-1] = ""     // Удалить последний элемент (записать нулевое значение)
	b = b[:len(b)-1]     // Усечь срез

	fmt.Println(b) // [A B D E]
}