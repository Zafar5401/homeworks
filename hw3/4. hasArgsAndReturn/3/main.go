package main

import "fmt"

func main() {
	fmt.Println(IsEven(1), IsEven(200))
}

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
