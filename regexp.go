package main

import (
	"fmt"
	"regexp"
)

func main() {

	/*
		3 ways to use regexp
		---
		1. MustCompile - panic kalau regex invalid - pattern statis, global, high performance
		2. Compile - return error - pattern dinamis (user input)
		3. MatchString - one-shot matching - cek cepat sekali saja

	*/
	var validateEmail *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	fmt.Println(validateEmail.MatchString("User@mail.com"))

	var findNums *regexp.Regexp = regexp.MustCompile(`\b([a-z]+):\s*([0-9]+)`)
	fmt.Println(findNums.FindAllString("id: 1233; usename: user3434; phone: 12131", -1))

	// ---

	ext := "png|jpg|jpeg"
	re, err := regexp.Compile(fmt.Sprintf(`\.(%s)$`, ext))
	// re, err := regexp.Compile(`\k`) // error since the pattern is invalid

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(re.MatchString("avatar.png")) // true
	fmt.Println(re.MatchString("file.txt"))   // false

	// ---

	var username string = "user1213"
	mustContainNumber, _ := regexp.MatchString(`[0-9]`, username)
	fmt.Printf("username %s contain number: %t \n", username, mustContainNumber)

}
