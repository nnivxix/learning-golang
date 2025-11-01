package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		// nil is default value or can assign to data type => function, map, slice, channel and pointer
		return nil
	}
	return map[string]string{
		"name": name,
	}
}
func main() {
	data := newMap("test")

	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}
