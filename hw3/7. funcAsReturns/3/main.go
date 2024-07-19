package main

import "fmt"

func ConditionalPrinter(condition bool) func(string) {
	return func(message string) {
		if condition {
			fmt.Println(message)
		}
	}
}

func main() {
	printIfTrue := ConditionalPrinter(true)

	printIfTrue("This will be printed")

	printIfFalse := ConditionalPrinter(false)
	printIfFalse("This will not be printed")
}
