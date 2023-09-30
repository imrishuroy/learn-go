package main

import "fmt"

func main() {

	fmt.Println("Welcome to structs")
	// no inheritance in golang; No super or parent

	rishu := User{"Rishu", "rishu@gmail.com", true, 21}
	fmt.Println(rishu)
	fmt.Printf("rishu details are: %+v\n", rishu)

	rishu.GetStatus()
	rishu.NewMail()
	fmt.Printf("rishu details are: %+v\n", rishu)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// when we pass a struct object is passed as copy
// so if we want to change the value of struct we need to pass it as pointer

func (u User) NewMail() {
	u.Email = "test@go.dev"

	fmt.Println("Email of this user is: ", u.Email)
}
