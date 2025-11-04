package main

import (
	"errors"
	"fmt"
)

func SendMessage(message string) (string, error) {
	if message == "" {
		return message, errors.New("message cannot blank")
	}
	return message, nil
}

func main() {
	message, err := SendMessage("")

	if err == nil {
		fmt.Println("sent =>", message)
	} else {
		fmt.Println("error =>", err.Error())
	}
}
