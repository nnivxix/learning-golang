package main

import "fmt"

type Blacklist func(name string) bool

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	register("anjing", blacklist)
	register("Hanasa", func(name string) bool {
		return name == "anjing"
	})
}

func register(name string, filter Blacklist) {
	if filter(name) {
		fmt.Println("You are blocked:", name)
	} else {
		fmt.Println("Welcome: ", name)
	}
}
