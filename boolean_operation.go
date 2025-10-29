package main

import "fmt"

func main() {
	var finalResult = 80
	var presentResult = 75

	var isPassedFinal = finalResult > 80
	var isPassedPresent = presentResult >= 75

	var isPassed = isPassedFinal && isPassedPresent

	// var isPassed = presentResult >= finalResult
	fmt.Println("Has the student passed?", isPassed)
}
