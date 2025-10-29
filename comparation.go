package main

import "fmt"

func main() {

	var name1 = "Alice"
	var name2 = "Bob"

	var result = name1 == name2
	fmt.Println("Are the names equal?", result)

	var age1 = 25
	var age2 = 30

	var isOlder = age1 > age2
	fmt.Println("Is age1 greater than age2?", isOlder)
}
