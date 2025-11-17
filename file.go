package main

import (
	"bufio"
	"io"
	"os"
)

const FILENAME = "sample.log"

func main() {
	// createFile(FILENAME, "this sample log\n")

	// content, _ := readFile(FILENAME)
	// fmt.Println(content)

	updateFile(FILENAME, "\nNew log")

}

func createFile(name string, content string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(content)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var content string
	for {
		line, _, err := reader.ReadLine()
		content += string(line) + "\n"

		if err == io.EOF {
			break
		}
	}
	return content, nil
}

func updateFile(name string, content string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)

	return nil
}
