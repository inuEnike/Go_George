package main

import "fmt"

func main() {
	Bio("george", 20, "Nigeria", "Male")
}

func Bio(name string, age int, nationality string, sex string) {
	sayName := name
	sayAge := age
	sayNationality := nationality
	saySex := sex

	fmt.Println("My name is", sayName, "\n", "and I am from", sayNationality, "\n", "I am", sayAge, "years old", "\n", "I am a", saySex)
}
