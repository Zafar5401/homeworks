package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(average(10, 5))
}

func average(a, b float64) (avg float64) {
	if a < 0.9 || b < 0.9 {
		fmt.Println("Нельзя передавать отрицательное число или число близкое к нулю!")
		return
	}
	avg = math.Sqrt(a * b)
	return
}
