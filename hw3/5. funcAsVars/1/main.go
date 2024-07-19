package main

import "fmt"

func main() {
	adder := func(a, b int) int {
		return a + b
	}

	fmt.Println(adder(2, 3))
}
