package main

import "fmt"

func main() {
	fmt.Println(Concat("Some ", "text"))
}

func Concat(a, b string) string {
	return a + b
}
