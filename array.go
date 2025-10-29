package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	fmt.Println("Names:", names)

	var scores = [4]int{90, 85, 88, 92}

	fmt.Println("Scores index 1:", scores[1])

	var fruits = [...]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Have fruits:", len(fruits))

}
