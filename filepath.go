package main

import (
	"fmt"
	"path/filepath"
)

// normally case when handle with OS
func main() {

	fmt.Println(filepath.Base("/a/b/c.txt")) // return last element
	fmt.Println(filepath.Ext("/a/b/c.txt"))  // return .txt

	// Split - pisah dir dan file
	dir, file := filepath.Split("/home/user/doc")
	fmt.Printf("Dir: %s, File: %s\n", dir, file)

	fmt.Println(filepath.Join("home", "user", "Downloads"))

}
