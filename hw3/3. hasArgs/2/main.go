package main

import "fmt"

func main() {
	GreetUser("user")
}

func GreetUser(user string) {
	fmt.Println("Hello, " + user)
}
