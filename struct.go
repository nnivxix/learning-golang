package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

// struct method
func (user User) sayHello(name string) {
	fmt.Println("Hello ", name, "My name is ", user.Name)
}

func main() {

	var leon User

	leon.Name = "Leon"
	leon.Age = 29
	leon.IsActive = false

	fmt.Println(leon)
	fmt.Println("name:", leon.Name)

	ada := User{
		Name:     "Ada",
		Age:      27,
		IsActive: true,
	}

	fmt.Println(ada)

	luis := User{"Luis", 30, true}

	fmt.Println(luis)

	// struct method
	luis.sayHello("Wesker")
}
