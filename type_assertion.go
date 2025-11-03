package main

import "fmt"

func main() {
	var result any = message()
	// fmt.Println(result)

	// var strResult string = result.(string)
	// fmt.Println(strResult)

	// using swicth is more safe than force to direct assert
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Default", value)
	}

}

func message() any {
	return 1213
}
