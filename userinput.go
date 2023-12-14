package main

import "fmt"

// This is my first Go lang program
func main() {

	var name string
	var food bool

	fmt.Print("Enter Your name & Have you Eaten? ")
	fmt.Scanf("%s %t ", &name, &food)

	fmt.Println(name, food)
}
