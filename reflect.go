package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Name string `required:"true" max:"7"`
	Age  int    `max:"26"`
}

func main() {
	user1 := User{"Hanasa", 25}

	user1Type := reflect.TypeOf(user1)   // get type
	user1Value := reflect.ValueOf(user1) // get value

	for i := 0; i < user1Type.NumField(); i++ {
		typeOf := user1Type.Field(i)
		value := user1Value.Field(i)

		if typeOf.Name == "Age" && typeOf.Type.Kind() == reflect.Int {
			maxAge := typeOf.Tag.Get("max")

			if int(value.Int()) > atoi(maxAge) {
				fmt.Println(errors.New("age to old").Error())
			} else {
				fmt.Println("you can login now")
			}

		}
	}

	example := 42
	reflectType := reflect.TypeOf(example)
	fmt.Println(reflectType) // int
	reflectValue := reflect.ValueOf(example)
	fmt.Println(reflectValue.Kind()) // int

	var emptyVal interface{} = nil
	vInvalid := reflect.ValueOf(emptyVal)
	fmt.Println("emptyVal:", vInvalid, "Kind:", vInvalid.Kind(), "IsValid:", vInvalid.IsValid())
	// if emptyValReflect.IsNil() {
	// 	fmt.Println("is nil", emptyValReflect.IsNil())
	// }

	// var ptrNil *User = nil
	// var emptyIface interface{} = ptrNil
	// vPtrNil := reflect.ValueOf(emptyIface)
	// fmt.Println("emptyIface:", vPtrNil, "Kind:", vPtrNil.Kind(), "IsValid:", vPtrNil.IsValid(), "IsNil:", vPtrNil.IsNil())

	// var userNil *User = nil
	// var

}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)

	return val
}
