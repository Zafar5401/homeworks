package main

import "fmt"

func main() {
	c := summary(5, 3)
	fmt.Println(c)

	c = difference(5, 3)
	fmt.Println(c)

	c = product(5, 3)
	fmt.Println(c)

	c = quotient(5, 3)
	fmt.Println(c)
}

func summary(a, b float64) (c float64) {
	c = a + b
	return
}

func difference(a, b float64) (c float64) {
	c = a - b
	return
}

func product(a, b float64) (c float64) {
	c = a * b
	return
}

func quotient(a, b float64) (c float64) {
	c = a / b
	return
}
