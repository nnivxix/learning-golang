package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Task 1")
	data.PushBack("Task 2")
	data.PushBack("Task 3")
	data.PushBack("Task 4")

	// fmt.Println(data.Front().Value)
	// data.Remove(data.Front())

	// for e := data.Front(); e != nil; e = e.Next() {
	// 	// fmt.Printf("the value %s is from type %T \n", e.Value, e.Value)
	// 	fmt.Printf("Do %s\n", e.Value)
	// 	// data.Remove(e)
	// }

	for data.Len() > 0 {
		e := data.Front()
		fmt.Printf("Do %s\n", e.Value)
		data.Remove(e)

	}
	fmt.Printf("Remain tasks (%v)\nYou can continue working on another task\nThank you\n", data.Len())
}
