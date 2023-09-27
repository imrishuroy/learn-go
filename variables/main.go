package main

import "fmt"

const LoginToken string = "asdasd" // this is a public constant variable

func main() {

	fmt.Println("Hi there")

	// string
	var username string = "rishu"
	fmt.Println(username)

	fmt.Printf("Variable is of type: %T \n", username)

	// boolean
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// uint8 (0-255)
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// flotat32 - this will give us only decimal digits upto 5 values
	var smallFloat float32 = 255.97972121221571
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// flotat32 - this will give us only decimal digits upto 5 values
	var bigFloat float64 = 255.97972121221571
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	// by default we have 0, if we only declare a int variable and
	// haven't initialized it
	var anotherInVariable int
	fmt.Println(anotherInVariable)
	fmt.Printf("Variable is of type: %T \n", anotherInVariable)

	// by default we have empty string, if we only declare a string variable
	// and haven't initialized it
	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable is of type: %T \n", anotherStringVariable)

	// implicit tye
	var website = "github.com"
	fmt.Println(website)
	// the website variable is now declared as string
	// we can't do website = 20

	// no var style
	numberOfUsers := 300000 // we can't use this outside a function
	fmt.Println(numberOfUsers)

	fmt.Println("Login Token: ", LoginToken)
	fmt.Printf("Login Token is of type: %T \n", LoginToken)

}
