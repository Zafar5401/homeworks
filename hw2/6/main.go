package main

import "fmt"

func main() {
	fmt.Println(volume(5, 6, 1.5))
	fmt.Println(square(5, 6, 1.5))
}

func volume(a, b, c float64) (v float64) {
	v = a * b * c
	return
}

func square(a, b, c float64) (s float64) {
	s = 2 * (a*b + b*c + a*c)
	return
}
