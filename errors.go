package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("not found error")
)

func main() {
	result, err := GetById(-1)

	if err != nil {
		// fmt.Println(err.Error())
		if errors.Is(err, validationError) {
			fmt.Println(validationError.Error())
		} else if errors.Is(err, notFoundError) {
			fmt.Println(notFoundError.Error())
		} else {
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println(result)
	}

	joinErrors := errors.Join(validationError, notFoundError)
	fmt.Println(joinErrors)

	if errors.As(joinErrors, &err) {
		fmt.Println("pasti not found")
	}

	level1 := errors.New("network timeout")
	// %w wrapping error
	level2 := fmt.Errorf("connection failed: %w", level1)
	level3 := fmt.Errorf("service unavailable: %w", level2)

	fmt.Println(errors.Unwrap(level3))
}

func GetById(id int) (string, error) {
	if id == 0 {
		return "", validationError
	} else if id < 0 {
		return "", notFoundError
	} else {
		return fmt.Sprintf("id found %d", id), nil
	}
}
