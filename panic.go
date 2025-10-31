package main

import "fmt"

func main() {
	defer logging() // will invoked at end of program/function even program error
	if isError(true) {
		panic("awokwowk error")
	}

}

func logging() {
	fmt.Println("logging")
}

func isError(state bool) bool {
	return state
}
