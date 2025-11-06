package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("xa halo axxx", "x")) //a halo a
	fmt.Println(strings.ToUpper("helo"))
	fmt.Println(strings.ToLower("HELO"))
	fmt.Println(strings.Contains("hello", "he"))    // true
	fmt.Println(strings.ContainsAny("hello", "eq")) // true

	// cut
	str := "halo selamat pagi"
	before, after, found := strings.Cut(str, "lo")
	fmt.Printf("cut string from '%s' using 'lo' return before: %s, after: %s, found: %t \n", str, before, after, found)

	fmt.Println(strings.Split("halo semuanya", " "))
}
