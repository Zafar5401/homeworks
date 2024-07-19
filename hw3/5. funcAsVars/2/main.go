package main

import "fmt"

func main() {
	concatenator := func(a, b string) string {
		return a + b
	}

	fmt.Println(concatenator("Hello, ", "world!"))
}
