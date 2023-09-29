package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to switch case")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 2:
		fmt.Println("You can move to 2 spot")
		fallthrough
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough // to run the next case also
	case 4:
		fmt.Println("You can move to 4 spot")

	case 5:
		fmt.Println("You can move to 5 spot")

	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")

	default:
		fmt.Println("What was that!")
	}

}
