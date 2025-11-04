package main

import "fmt"

type User struct {
	id   int
	name string
	age  int
}

func main() {
	// var user1 *User = &User{} // return empty data
	// var user2 = user1

	// user2.id = 1

	// you can use new

	var user1 *User = new(User) // return empty data like above code
	var user2 *User = user1

	user1.name = "Mia"
	user2.name = "Robert"

	fmt.Println(user1)
	fmt.Println(user2)
}
