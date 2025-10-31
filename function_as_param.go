package main

import "fmt"

type Filter func(string) string

func main() {
	sayHello("Hanasa", spamFilter)
	sayHello("Anjing", spamFilter)

}

func sayHello(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	}
	return name
}
