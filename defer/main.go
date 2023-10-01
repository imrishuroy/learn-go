package main

import "fmt"

func main() {
	// last in first out

	// fmt.Println("Hello")
	// defer fmt.Println("World")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello") // this will run first

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
