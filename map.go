package main

import "fmt"

func main() {
	var user = map[string]string{
		"name":    "Alice",
		"country": "Wonderland",
		"job":     "Adventurer",
	}
	fmt.Println("user name:", user["name"])
	fmt.Println("user hobby:", user["hobby"]) // Accessing a non-existing key will return empty string
	fmt.Println("User Map:", user)

	delete(user, "job")

	fmt.Println("After deleting job:", user)

	book := make(map[string]string)
	book["title"] = "Go Programming"
	book["author"] = "John Doe"

	fmt.Println("Book Map:", book)
}
