package main

import "fmt"

func main() {
	fmt.Println(Apply(2, func(a int) int {
		return a * a
	}))
}

func Apply(a int, f func(int) int) int {
	return f(a)
}
