package main

import "fmt"

type User struct {
	id   int
	name string
	age  int
}

func main() {
	// var user1 User = User{1, "Hanasa", 27}
	// var user2 User = user1 // copied value from user1

	// user2.id = 2

	// fmt.Println(user1) // values will keep same
	// fmt.Println(user2)

	// & operator to make value inheritance

	var user1 = User{1, "Hanasa", 27}
	var user2 = &user1 // inheritance value from user1

	user2.id = 2 // user1 value will change as well

	fmt.Println(user1) // values will keep same
	fmt.Println(user2)

}
