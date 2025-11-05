package main

import (
	"fmt"
	myfmt "go-standard-lib/myfmt"
)

func main() {
	// name := "Hanasa"
	// age := 27

	// var contents = make([]string, 3)

	text := "Hello"
	textBytes := []byte(text)
	fmt.Printf("String '%s' sebagai bytes: %v\n", text, textBytes)

	// fmt.Printf("Hallo my name is %s I'm %d years old \n", name, age)
	// fmt.Print("Hallo my name is ", name, " I'm ", age, " years old \n")
	// // saving string value
	// message := fmt.Sprintln("Hallo my name is ", name, " I'm ", age, " years old ss \n")
	// // then print the value
	// fmt.Println(message)

	myfmt.Append()
	myfmt.GetError()

}
