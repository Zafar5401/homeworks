package main

import "fmt"

func main() {
	fmt.Println(square(8))
}

func square(a float64) (s float64) {
	s = a * a
	return
}
