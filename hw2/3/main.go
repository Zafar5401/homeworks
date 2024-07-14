package main

import "fmt"

func main() {
	fmt.Println(square(5, 4))
	fmt.Println(perimeter(5, 4))
}

func square(a, b float64) (s float64) {
	s = a * b
	return
}

func perimeter(a, b float64) (p float64) {
	p = 2 * (a + b)
	return
}
