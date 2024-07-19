package main

import "fmt"

func main() {
	GetGrade()
}

func GetGrade() {
	var score int8
	fmt.Scan(&score)
	switch score {
	case 5:
		fmt.Println("A")
	case 4:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	}
}
