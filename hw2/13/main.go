package main

import "fmt"

func main() {
	fmt.Println(byteToKb(16384))
}

func byteToKb(b float64) (kb float64) {
	kb = b / 1024
	return
}
