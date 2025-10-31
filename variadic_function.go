package main

import "fmt"

func main() {
	// variable arguments
	fmt.Println(sumAll(1, 2, 3, 4, 5))

	var numbers = []int{1, 2, 3, 4, 5}
	// using slice
	fmt.Println(sumAll(numbers...))
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
