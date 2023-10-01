package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	//PerformGetRequest()
	//PerformJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code ", response.StatusCode)
	fmt.Println("content length ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"name": "Rishu",
			"age": 21
		}
	`)

	response, error := http.Post(myUrl, "application/json", requestBody)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("firstname", "rishu")
	data.Add("lastname", "kumar")
	data.Add("email", "rishu@gmail.com")

	response, error := http.PostForm(myUrl, data)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
