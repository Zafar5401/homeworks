package main

import (
	"fmt"
	"strings"
)

func main() {
	repeat3Times := StringRepeater(3)
	result := repeat3Times("Zafar")
	fmt.Println(result)
}

func StringRepeater(n int) func(string) string {
	return func(s string) string {
		return strings.Repeat(s, n)
	}
}
