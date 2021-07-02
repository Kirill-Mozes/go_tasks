//созать мапу int clice int где ключ это чило /10
package main

import "fmt"

func main() {
	var temperatues = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0.3, 1.3, -1.1, 0.0}
	m := make(map[int][]float64)
	for _, v := range temperatues {
		key := int(v / 10) //возникает проблема с числами от -10 до 10
		if key == 0 && v >= 0 {
			key = 1 // [0, 10)
			m[key] = append(m[key], v)
		} else if key == 0 && v < 0 {
			key = -1 // (-10, 0)
			m[key] = append(m[key], v)
		} else {
			key *= 10
			m[key] = append(m[key], v)
		}
	}
	fmt.Println(m)

}
