package main

import "fmt"

func main() {
	GetTemperatureDescription()
}

func GetTemperatureDescription() {
	var temperature int
	fmt.Scan(&temperature)
	switch {
	case temperature < 10:
		fmt.Println("Холодно")
	case temperature >= 10 || temperature <= 25:
		fmt.Println("Тепло")
	default:
		fmt.Println("Жарко")
	}
}
