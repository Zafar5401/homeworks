package main

import "fmt"

func main() {
	isPositive := func(a int) bool {
		if a >= 0 {
			return true
		} else {
			return false
		}
	}

	fmt.Println(isPositive(25), isPositive(-11), isPositive(0))

}
