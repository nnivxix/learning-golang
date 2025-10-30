package main

import "fmt"

func main() {
	firstName, _, _ := getFullName()

	fmt.Println(firstName)
}

func getFullName() (firstName, middleName, lastName string) {

	firstName = "Ade"
	middleName = "Hanasa"
	lastName = "Sofari"

	return firstName, middleName, lastName
}
