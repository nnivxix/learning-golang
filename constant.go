package main

import "fmt"

func main() {
	const FILENAME string = "data.txt"
	const MAXRETRY int = 5
	const PI float64 = 3.14159

	// FILENAME = "config.txt" // This will cause a compile-time error

	fmt.Println("Filename:", FILENAME)
	fmt.Println("Max Retry:", MAXRETRY)
	fmt.Println("PI:", PI)

	fmt.Println("----")

	const (
		STATUS_OK       = 200
		STATUS_ERROR    = 500
		STATUS_NOTFOUND = 404
	)

	fmt.Println("Status OK:", STATUS_OK)
	fmt.Println("Status Error:", STATUS_ERROR)
	fmt.Println("Status Not Found:", STATUS_NOTFOUND)
}
