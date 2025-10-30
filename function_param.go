package main

import "fmt"

func main() {
	hello("Hanasa")
	calculate(2, 3)
}

func hello(name string) {
	fmt.Println("hello ", name)
}

func calculate(x int8, y int8) {
	fmt.Println("result :", x+y)
}
