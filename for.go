package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	println("Counter:", counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	println("Counter:", counter)
	// }

	println("Finished!")

	names := []string{"Leon", "Ada", "Ashley"}

	// manual
	for i := 0; i < len(names); i++ {
		fmt.Println("manual looping => ", names[i])
	}

	// using range
	for key, value := range names {
		fmt.Println("key: ", key, "value:", value)
	}

	// but if you dont need a key you can use _
	for _, value := range names {
		fmt.Println(value)
	}
}
