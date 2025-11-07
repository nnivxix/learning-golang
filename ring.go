package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// assum will creating header table
func main() {
	// Create a new ring of size 5
	headers := ring.New(5)

	// Initialize the ring with some values
	for i := 0; i < headers.Len(); i++ {
		// fill value
		headers.Value = "Header " + strconv.Itoa(i+1)
		// reassign again
		headers = headers.Next()
	}

	// Iterate through the ring and print its contents
	headers.Do(func(h any) {
		fmt.Println(h)
	})
}
