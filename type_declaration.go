package main

import "fmt"

func main() {
	type uid string

	var userID uid = "user_12345"
	fmt.Println("User ID:", userID)

	var aliceId uid = uid("alice_67890")
	fmt.Println("Alice ID:", aliceId)
}
