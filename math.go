package main

import "fmt"

func main() {

	var a = 10
	var b = 5
	var c = 15

	sum := a + b + c
	fmt.Println("Sum:", sum)

	subtract := a - b - c
	fmt.Println("Subtract:", subtract)

	multiply := a*b + c
	fmt.Println("Multiply:", multiply)

	divide := c / b
	fmt.Println("Divide:", divide)

	modulus := c % a
	fmt.Println("Modulus:", modulus)

	// Unary operations
	negate := -a
	fmt.Println("Negate:", negate)

	increment := a
	increment++
	fmt.Println("Increment:", increment)

	a = a + 10
	fmt.Println("After adding 10, a:", a)

}
