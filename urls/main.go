package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?course=react&paymentid=asoh212"

func main() {
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) // params

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["course"])

	for _, value := range qparams {
		fmt.Println("Param is ", value)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=aman",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
