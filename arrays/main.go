package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Peach"
	fruits[3] = "Mango"

	fmt.Println("Fruits lists ", fruits)
	fmt.Println("Fruits lists ", len(fruits))

	var vegList = [3]string{"Patato", "beans", "mushroom"}
	fmt.Println("Vegy list ", vegList)

}
