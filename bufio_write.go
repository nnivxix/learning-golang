package main

import (
	"bufio"
	"os"
)

func main() {
	text := os.Args[1:] // take all args from index 1 since the first index[0] is memory localtion
	// [/var/folders/fy/8t97ww1j2gs54rkmx12x8v0r0000gn/T/go-build2211444366/b001/exe/bufio_write Halo dunia]
	writer := bufio.NewWriter(os.Stdout)
	for index, _ := range text {

		_, _ = writer.WriteString(text[index] + " ")
	}

	writer.Flush()

	//  go run bufio_write.go Halo dunia
	// Halo dunia %

}
