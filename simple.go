package main

import "fmt"

// This is my first Go lang program
func main() {

	var name string
	fmt.Print("Enter Your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hey there, ", name)
}
