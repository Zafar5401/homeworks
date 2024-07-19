package main

import "fmt"

func main() {
	PrintTrafficLight()
}

func PrintTrafficLight() {
	var lightColor string
	fmt.Scan(&lightColor)
	switch lightColor {
	case "красный":
		fmt.Println("Стоп")
	case "желтый":
		fmt.Println("Внимание")
	case "зеленый":
		fmt.Println("Идите")
	default:
		fmt.Println("Светофор сломался")
	}
}
