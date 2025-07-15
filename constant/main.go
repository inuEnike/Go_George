package main

import "fmt"

func main() {
	// Constant are immutable value that cannot be changed

	//  const Name string = "Inu Enike"
	//  const Age int = 20
	//  const Nationality = "Nigeria"

	const (
		Name        string  = "Inu Enike"
		Age         int     = 20
		Nationality string  = "Nigeria"
		Networth    float32 = 0.00
	)

	fmt.Println("Name:", Name, "\n", "Age:", Age, "\n", "Nationality:", Nationality, "Networth:", Networth)

	// Escape Sequence in string literals

	/**
		New Line: \n
		Quote "\"Go Is simple\" - A Programmer"
		Line Break: "I\nAM\nBATMAN\n"
	**/

}
