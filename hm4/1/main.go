package main

import "fmt"

func main() {
	PrintGreeting()
}

func PrintGreeting() {
	var dayType string
	fmt.Scan(&dayType)
	switch dayType {
	case "утро":
		fmt.Println("Доброе утро!")
	case "день":
		fmt.Println("Добрый день!")
	case "вечер":
		fmt.Println("Добрый вечер!")
	default:
		fmt.Println("Неизвестное время!")
	}
}
