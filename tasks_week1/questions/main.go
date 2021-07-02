package main

func main() {
	bs := make([]byte, 10)
	bl := 0
	for n := 0; n < 10; n++ {
		bl += copy(bs[bl:], "x")
	}
}
