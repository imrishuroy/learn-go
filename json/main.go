package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`              // I don't want this to be reflected in my api
	Tags     []string `json:"tags,omitempty"` // if the list is nill, the field will not appear in JSON
}

func main() {

	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {

	courses := []course{
		{"Flutter Dev", 400, "udemy.com", "abc123", []string{"mobile", "flutter"}},
		{"Golang", 450, "udemy.com", "abc100", []string{"go", "backend"}},
		{"Android", 400, "udemy.com", "abc123", nil},
	}

	// package this data as JSON data

	// finalJson, err := json.Marshal(courses)

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonData := []byte(`
	
	{
		"coursename": "Flutter Dev",
		"price": 400,
		"platform": "udemy.com",
		"tags": ["mobile", "flutter"]
    }
	
	`)

	var myCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON validated")
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// convert this json data into key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf(" key is %v and value is %v\n", k, v)
	}

}
