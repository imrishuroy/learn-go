package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {

	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // close the connection

	data, error := io.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	content := string(data)

	fmt.Println(content)

}
