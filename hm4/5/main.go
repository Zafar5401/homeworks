package main

import "fmt"

func main() {
	GetDiscount()
}

func GetDiscount() {
	var amount int
	fmt.Scan(&amount)
	switch {
	case amount > 1000:
		fmt.Println("10%")
	case 500 < amount || amount >= 100:
		fmt.Println("5%")
	case amount < 100:
		fmt.Println("0%")
	}
}
