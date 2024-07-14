package main

import "fmt"

func main() {
	fmt.Println(perimeter(4))
}

func perimeter(a float64) (p float64) {
	p = a * 4
	return
}
