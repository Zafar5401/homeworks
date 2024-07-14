package main

import "fmt"

func main() {
	fmt.Println(massToTons(200000))
}

func massToTons(m float64) (t float64) {
	t = m / 10000
	return
}
