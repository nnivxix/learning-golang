package main

import "fmt"

func main() {

	var score = 90

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// short statement
	if value := -2; value >= 0 {
		fmt.Println(value, "is a positive number")
	} else {
		fmt.Println(value, "is a negative number")
	}
}
