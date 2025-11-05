package myfmt

import "fmt"

func Append() {

	name := "Andi"
	// ✅ Efisien - menggunakan buffer
	var buf []byte
	buf = fmt.Appendf(buf, "Hello %s!", name)
	result := string(buf)

	fmt.Println(result)

	// Web server response
	var response []byte
	response = fmt.Append(response, "HTTP/1.1 200 OK\n")
	response = fmt.Append(response, "Content-Type: text/html\n\n")
	response = fmt.Append(response, "<h1>Hello %s</h1>", name)

	// convert byte to string
	fmt.Println(string(response))
}
