package main

import "fmt"

type User struct {
	id   int
	name string
	age  int
}

func main() {
	// var score int = 1

	// var mia_score = &score

	// fmt.Println(score)
	// fmt.Println(mia_score)

	// var user1 = User{1, "Hanasa", 27}
	// var user2 = &user1 // inheritance value from user1

	// user2.id = 2                // user1 value will change as well
	// user2 = &User{3, "Mia", 26} // REASSIGNMENT - new pointer

	// fmt.Println(user1)
	// fmt.Println(user2)

	// * (Asterisk) Op

	var user1 = User{1, "Hanasa", 27}
	var user2 = &user1 // inheritance value from user1

	user2.id = 2 // user1 value will change as well
	// * Operator
	*user2 = User{3, "Mia", 26} // REASSIGNMENT - change pointer value

	fmt.Println(user1)
	fmt.Println(user2)

}
