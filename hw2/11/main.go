package main

import "fmt"

func main() {
	fmt.Println(cmToMeters(1000))
}

func cmToMeters(cm float64) (m float64) {
	m = cm / 100
	return
}
