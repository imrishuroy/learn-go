package main

import "fmt"

func main() {

	// In the Go programming language, a pointer is a variable that stores
	// the memory address of another value, rather than the value itself.
	// Pointers are often used to indirectly access and manipulate data stored in memory

	fmt.Println("Welcome to pointers")

	// var ptr *int
	// * is for actual value
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	// & is used for memory ref of variable
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2 // pointer makes it a guarantee,
	// no matter what you are doing operation on that values
	// those operations are actually perform on the value on the copy of value
	fmt.Println("New value is ", myNumber)

}
