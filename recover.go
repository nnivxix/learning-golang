package main

import "fmt"

func main() {
	defer endApp()
	if isError(false) {
		panic("Upp error")
	}
}

func endApp() {
	fmt.Println("app ended")
	message := recover()
	if message != nil {
		fmt.Println("whats error =>", message)
	}
}

func isError(state bool) bool {
	return state
}
