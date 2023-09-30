package main

import "fmt"

func main() {

	fmt.Println("Welcome to functions")

	greet()

	result := add(2, 5)

	fmt.Println(result)

	proResult, message := proAdder(2, 3, 2, 5, 1)

	fmt.Println("Pro result ", proResult, " and message is ", message)
}

func greet() {
	fmt.Println("Namstey from golang")
}

func add(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {

	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi pro result functions"
}
