package main

import "fmt"

func main() {
	Execute(true, func(b bool) {
		b = !b
		fmt.Println(b)
	})
}

func Execute(b bool, f func(bool)) {
	f(b)
}
