package main

import "fmt"

func main() {
	sapa := hello

	fmt.Println(sapa("Hanasa"))
}

func hello(name string) string {
	return "Hello " + name
}
