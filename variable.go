package main

import "fmt"

func main() {
	var name string

	name = "Alice"
	fmt.Println("Name:", name)

	name = "Bob"
	fmt.Println("Updated Name:", name)

	fmt.Println("----")

	// first initialize can using :=
	age := "30"
	fmt.Println("Age:", age)

	// but when updating must use =
	age = "31"
	fmt.Println("Updated Age:", age)

	fmt.Println("----")

	var (
		firstName string = "Charlie"
		lastName  string = "Brown"
	)
	fullName := firstName + " " + lastName
	fmt.Println("Full Name:", fullName)
}
