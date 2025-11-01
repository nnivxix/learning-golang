package main

import "fmt"

func main() {
	fmt.Println(getValue())
}

// any on interface{}
func getValue() interface{} {
	// return "hello"
	// return true
	return 1
}
