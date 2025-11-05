package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)

		/*
			input: go run os.go halo hanasa
			/----\
			output:
			/var/folders/fy/8t97ww1j2gs54rkmx12x8v0r0000gn/T/go-build1083618047/b001/exe/os
			halo
			hanasa
		*/
	}

	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("error, ", err.Error())
	} else {
		fmt.Println(hostname)
	}

	err = os.Mkdir("test", 0750)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	files, readErr := os.ReadDir("myfmt")

	if readErr != nil {
		fmt.Println("error", readErr.Error())
	}

	for _, file := range files {
		fmt.Println(file)
	}

	writeError := os.WriteFile("text.txt", []byte("halo"), 0666)

	if writeError != nil {
		fmt.Println("error write file", writeError.Error())
	}

	readFile, errorRead := os.ReadFile("text.txt")
	if errorRead != nil {
		fmt.Println("error", errorRead.Error())
	}
	fmt.Println(string(readFile))
	// os.Stdout.Write(readFile)

}
