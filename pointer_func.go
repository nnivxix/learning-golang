package main

import "fmt"

type User struct {
	id   int
	name string
	age  int
}

// since pointer using asterisk, the value will change
func changeUser(user *User) {
	user.name = "Hanasa"
}
func main() {
	user1 := User{}

	changeUser(&user1)

	fmt.Println(user1)
}
