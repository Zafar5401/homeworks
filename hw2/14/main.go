package main

import "fmt"

func main() {
	fmt.Println(cutAmount(100, 20))
}

func cutAmount(a, b float64) (count float64) {
	if a > b {
		count = a / b
	} else {
		fmt.Println("B больше чем A!")
	}
	return
}
