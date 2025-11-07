package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2020, time.March, 5, 8, 8, 12, 0, time.Local)
	fmt.Println(utc)
	// fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2020-03-05T08:08:09Z")
	fmt.Println(parse.Local())
}
