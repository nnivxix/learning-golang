package main

import "fmt"

func main() {
	fmt.Println(5 * 4 * 3 * 2 * 1)

	fmt.Println(factorialLoop(5))
	fmt.Println(recursiveFunction(5))
}

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func recursiveFunction(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFunction(value-1)
	}
}
