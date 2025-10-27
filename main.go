package main

import "fmt"

func main() {

	fmt.Println("\n DATA TYPES AND VARIABLES \n")

	var number int = 10
	fmt.Println("The value of number is:", number)

	var message string = "Hello, World!"
	fmt.Println("The message is:", message)

	var isActive bool = true
	fmt.Println("Is active:", isActive)

	var pi float64 = 3.14
	fmt.Println("The value of pi is:", pi)

	fmt.Println("\n SHORT HAND VARIABLE DECLARATION \n")
 
	dynamicNumber := 1234
	fmt.Println("The dynamic number is:", dynamicNumber)

	dynamicMessage := "Dynamic Hello!"
	fmt.Println("The dynamic message is:", dynamicMessage)

	fmt.Println("\n MULTIPLE VARIABLE DECLARATION AND CONSTANTS \n")

	firtName, lastName := "John", "Doe"

	fullName := firtName + " " + lastName
	fmt.Println("Full name is:", fullName)

	firstNumber, secondNumber := 5, 10
	sum := firstNumber + secondNumber
	fmt.Println("The sum of", firstNumber, "and", secondNumber, "is:", sum)

	const MAX_LIMIT int = 100
	fmt.Println("The maximum limit is:", MAX_LIMIT)

	const VERSION_NAME string = "Beta 1.0"
	fmt.Println("The version name is:", VERSION_NAME)

	// This will cause an error because VERSION_NAME is a constant
	// VERSION_NAME = "Release 1.0" 

	fmt.Println("\n OPERATORS \n")

	var num1 int = 20
	var num2 int = 30

	num1 += 10
	fmt.Println("After incrementing, num1 is:", num1)

	num2 -= 5
	fmt.Println("After decrementing, num2 is:", num2)

	num1 *= 2
	fmt.Println("After multiplying, num1 is:", num1)

	num2 /= 5
	fmt.Println("After dividing, num2 is:", num2)

	num1 %= 3
	fmt.Println("After modulus, num1 is:", num1)


	fmt.Println("\n COMPARISON OPERATORS \n")

	fmt.Println("Is 1 and 2 equal?", 1 == 2)
	fmt.Println("Is 1 not equal to 2?", 1 != 2)

	fmt.Println("\n LOGICAL OPERATORS \n")

	var authenticated bool = true
	var isAdmin bool = false

	fmt.Println("Is the user authenticated?", !authenticated) // false

	if authenticated && isAdmin {
		fmt.Println("User is authenticated and is an admin.")
	} else {
		fmt.Println("User is either not authenticated or not an admin.")
	}
}
