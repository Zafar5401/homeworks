package main

import "fmt"

func main() {
	fmt.Println(remainderAmount(100, 20))
}

func remainderAmount(a, b int64) (count int64) {
	if a > b {
		count = a % b
	} else {
		fmt.Println("B больше чем A!")
	}
	return
}
