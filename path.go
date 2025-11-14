package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/a/b/c.txt")) // return last element (c.txt)
	fmt.Println(path.Base("/a/b"))       // return last element (b)

	fmt.Println(path.Ext("/a/b/c.txt")) // return .txt
	fmt.Println(path.Ext("/a/b/c"))     // return empty string

	slug := "xsde"
	fmt.Println(path.Join("/api/v1", "post", slug))
}
