package main

import "fmt"

func main() {
	var mb int32 = 34000      // can store from -32,768 to 32,767
	var mb2 int16 = int16(mb) // explicit conversion from int32 to int16

	fmt.Println("Original MB value (int32):", mb)
	fmt.Println("Converted MB value (int16):", mb2)

	fmt.Println("----")

	name := "Golang"
	nameLength := len(name)                // get length of string
	firstChar := name[0]                   // get first character (byte)
	firstCharAsString := string(firstChar) // convert byte to string

	fmt.Println("Name Length:", nameLength)
	fmt.Println("First Character (byte):", firstChar)
	fmt.Println("First Character (string):", firstCharAsString)

}
