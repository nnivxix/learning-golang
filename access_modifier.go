package main

import (
	"basic/config"
	"fmt"
)

func main() {
	version := config.GetVersion()
	app_name := config.GetName()

	fmt.Println(version)
	fmt.Println(app_name)
}
