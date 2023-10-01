package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")

	fileName := "./myfile.txt"

	content := "This needs to go in a file"

	writeFile(fileName, content)

	readFile(fileName)

}

func writeFile(fileName string, content string) {
	file, error := os.Create(fileName)

	checkNilErr(error)

	length, error := io.WriteString(file, content)

	checkNilErr(error)

	fmt.Println("length is: ", length)
	defer file.Close()

}

func readFile(fileName string) {
	databyte, error := os.ReadFile(fileName)

	checkNilErr(error)

	fmt.Println("Text data inside the file is", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
