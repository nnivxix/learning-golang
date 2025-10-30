package main

import "fmt"

func main() {
	// firstName, middleName, lastName := fullName()
	// fmt.Println(firstName, middleName, lastName)

	firstName, _, _ := fullName()

	fmt.Println(firstName)
}

func fullName() (string, string, string) {
	return "Ade", "Hanasa", "Sofari"
}
