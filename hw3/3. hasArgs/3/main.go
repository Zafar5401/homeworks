package main

import "fmt"

func main() {
	ToggleLight(true)
	ToggleLight(false)
}

func ToggleLight(l bool) {
	if l == true {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}
