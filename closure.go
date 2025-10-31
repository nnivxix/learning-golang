package main

import "fmt"

func main() {
	//counter is ability to interact with other variables
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		// interact without parameter
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
