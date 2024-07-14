package main

import "fmt"

func main() {
	fmt.Println(average(10, 5))
}

func average(a, b float64) (avg float64) {
	avg = (a + b) / 2
	return
}
