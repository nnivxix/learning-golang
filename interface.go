package main

import "fmt"

type CanSpeak interface {
	Speak() string
}

type Animal struct {
	Name string
}

// implementation of interface inside struct
func (animal Animal) Speak(sound string) string {
	return animal.Name + " can speak:" + sound
}

func main() {
	kucing := Animal{"Kucing"}

	fmt.Println(kucing.Speak("meow"))
}
