package main

import "fmt"

func main() {

	fmt.Println(Calculate(1, 4, func(x, y int) int { return x + y }))
	fmt.Println(Calculate(2, 4, func(x, y int) int { return x * y }))
}

func Calculate(a, b int, f func(int, int) int) int {
	return f(a, b)
}
