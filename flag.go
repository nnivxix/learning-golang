package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "user", "name of user")
	var isAdmin *bool = flag.Bool("isAdmin", false, "set user wheater is admin or not")

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("is admin", *isAdmin)
}
