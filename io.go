package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	readFile, _ := os.Open("text.txt")

	defer readFile.Close()

	content, _ := io.ReadAll(readFile)
	fmt.Println(string(content))

}
