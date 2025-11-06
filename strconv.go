package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("1")
	fmt.Println(i, err)

	boo, booErr := strconv.ParseBool("0")
	fmt.Println(boo, booErr) // false <nil>

	formatBool := strconv.FormatBool(false)
	fmt.Printf("type %T , from value: %t \n", formatBool, false)
}
