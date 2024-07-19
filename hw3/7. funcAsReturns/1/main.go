package main

import "fmt"

func main() {
	multiplyBy3 := Multiplier(3)
	result := multiplyBy3(5)
	fmt.Println(result)
}

func Multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
