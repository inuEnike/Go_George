# LET'S GOOOOOOOOOOOOO

Started learning go today

### I've learnt

1. variable
2. Pointers
3. function
4. Data Types

### what is a Pointer

According to my own understanding, Pointers are variabe that store the address of another variable,
As a variable stores data into the memory location, Pointers allow access to the memory address of that data

### Functions

function is a block of code that can be reusable
eg instead of having to define names of 100 pupils in a class you can create a function that takes in a parameter of name, then you can call the functions anytime you want to process a pupil name

//main.go

package main

import "fmt"

func main() {

    Pupil("George")
    Pupil("Gregory")
    Pupil("John")
    Pupil("Trump")
    Pupil("Taiwo")
    Pupil("Kehinde")

}

func Pupil (name string) {

    pupilName := name
    fmt.Println(pupilName)

}
