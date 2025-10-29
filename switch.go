package main

import "fmt"

func main() {
	name := "Le"

	switch name {
	case "Leon":
		fmt.Println("Hello Leon!")
	case "Ada":
		fmt.Println("Hello Ada!")
	default:
		fmt.Println("Hello, who are you?")
	}

	// short switch statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Your name is long.")
	case false:
		fmt.Println("You can register your name.")
	}

	// switch with no expression
	switch {
	case len(name) < 3:
		fmt.Println("Your name is too short.")
	case len(name) > 10:
		fmt.Println("Your name is too long.")
	default:
		fmt.Println("Your name length is just right.")
	}
}
