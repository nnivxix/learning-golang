package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	usename := "Hanasa"
	var encoded = base64.StdEncoding.EncodeToString([]byte(usename))
	fmt.Println(encoded)

	var decode, _ = base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decode))
}
