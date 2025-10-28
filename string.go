package main

import "fmt"

func main() {
	var path string = "/user/local/bin"
	var filename string = "document.txt"
	var fullpath string = path + "/" + filename

	fmt.Println("Path:", path)
	fmt.Println("Filename:", filename)
	fmt.Println("Full Path:", fullpath)

	fmt.Println("----")

	fmt.Println("Length of Full Path:", len(filename))
	fmt.Println("First character of Filename:", filename[0])
}
