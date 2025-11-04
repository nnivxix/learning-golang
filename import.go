package main

// why using "basic"? because the module name using "basic"
// see go.mod
import (
	"basic/helper"
	"fmt"
)

func main() {
	message := helper.Hello("Han")
	fmt.Println(message)
}
