package myfmt

import "fmt"

func GetError() {
	var id int = 1

	err := fmt.Errorf("user id: %d not found", id)

	fmt.Println(err.Error())

}
