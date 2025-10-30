package main

import (
	"fmt"
)

func main() {
	message := hello("Hanasa")

	fmt.Println(message)
}

func hello(name string) string {
	return "hello " + name
}
