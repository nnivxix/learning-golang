package main

import "fmt"

type notFoundError struct {
	Message string
}

func (n notFoundError) Error() string {
	return n.Message
}

func FindById(id int) error {
	if id <= 0 {
		return &notFoundError{"User not found"}
	}
	return nil
}

func main() {
	id := 1
	err := FindById(id)

	if err != nil {
		fmt.Println("id =>", id, err.Error())
	} else {
		fmt.Println("data found with id =>", id)
	}
}
