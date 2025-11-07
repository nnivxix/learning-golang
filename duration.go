package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := time.Second * 53        // 53s
	duration2 := time.Millisecond * 3000 // 3s
	duration3 := duration1 - duration2
	m, _ := time.ParseDuration("1m30s")

	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Println(duration3) // 50s
	fmt.Println(m.Seconds())
}
