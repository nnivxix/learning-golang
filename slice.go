package main

import "fmt"

func main() {
	var fruits = []string{"Apple", "Banana", "Cherry", "Pineapple", "Mango", "Guava", "Papaya"}
	fmt.Println("Have fruits:", len(fruits))

	// var rujak = fruits[3:7]
	// var rujak = fruits[3:]
	var rujak []string = fruits[3:]
	// var rujak = fruits[:] // all elements

	fmt.Println("Rujak fruits:", rujak)
	fmt.Println("Rujak have fruits:", len(rujak))

	rujak = append(rujak, "Ambarella")
	fmt.Println("Rujak after append:", rujak)

	var salad = make([]string, 2, 5)
	salad[0] = "Watermelon"
	salad[1] = "Strawberry"

	fmt.Println("Salad fruits:", salad)
	fmt.Println("Salad length:", len(salad))
	fmt.Println("Salad capacity:", cap(salad))

	salad = append(salad, "Kiwi")
	salad = append(salad, "Grapes")
	salad = append(salad, "Orange")
	salad = append(salad, "Orangewood") // exceeds capacity

	fmt.Println("Salad after append:", salad)
	fmt.Println("Salad length after append:", len(salad))
	fmt.Println("Salad capacity after append:", cap(salad)) // capacity will double if exceeded

	juice := make([]string, len(salad), cap(salad))
	copy(juice, salad)
	fmt.Println("Juice fruits:", juice)

	thisArray := [3]string{"A", "B", "C"}
	thisSlice := []string{"D", "E", "F"}

	fmt.Println("This array:", thisArray)
	fmt.Println("This slice:", thisSlice)
}
