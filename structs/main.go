package main

import "fmt"

func main() {

	fmt.Println("Welcome to structs")
	// no inheritance in golang; No super or parent

	rishu := User{"Rishu", "rishu@gmail.com", true, 21}
	fmt.Println(rishu)
	fmt.Printf("rishu details are: %+v\n", rishu)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
