package main

import "fmt"

func main() {
	fmt.Println(lenght(12))
	fmt.Println(lenght(12))
}

var Pi = 3.14

func lenght(r float64) (l float64) {
	l = 2 * Pi * r
	return
}

func square(r float64) (s float64) {
	s = Pi * r
	return
}
