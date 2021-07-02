package main

import (
	"fmt"
	"math"
)

//Написать программу нахождения расстояния между 2 точками,
//которые представление в виде структуры Point с инкапсулированными
//параметрами x,y и конструктором.
type Point struct {
	x, y float64
}

func newPoint(x, y float64) Point {
	return Point{x, y}
}
func (first Point) distance(second Point) float64 {
	return math.Sqrt((first.x-second.x)*(first.x-second.x) + (first.y-second.y)*(first.y-second.y))
}
func main() {
	first := newPoint(1.1, 2.)
	second := newPoint(1., 1.)
	fmt.Println(first.distance(second))
}
