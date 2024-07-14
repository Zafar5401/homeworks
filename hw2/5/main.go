package main

import "fmt"

func main() {
	fmt.Println(volume(5.5))
	fmt.Println(square(5.5))
}

func volume(a float64) (v float64) {
	v = a * a * a
	return
}

func square(a float64) (s float64) {
	s = 6 * a
	return
}
