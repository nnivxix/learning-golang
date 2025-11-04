package main

import "fmt"

type RegistUser struct {
	name     string
	verified bool
}

// make sure to user * when using pointer to make value updated
func (user *RegistUser) verification() {
	user.verified = true
}
func main() {
	var user1 RegistUser = RegistUser{"Hanasa", false}

	user1.verification()

	fmt.Println(user1)
}
