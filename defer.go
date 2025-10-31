package main

import "fmt"

func main() {
	// will invoked at end of program/function
	defer logging()
	fmt.Println("App is starting...")
}

func logging() {
	fmt.Println("Logging")
}
