package main

import (
	"basic/database"
	_ "basic/internal" // using _ when package only has init function
	"fmt"
)

func main() {
	var database string = database.GetDatabase()

	fmt.Println(database) // mySql
}
