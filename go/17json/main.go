package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main()  {
	fmt.Println("Welcome to JSON video")

	// convert data to valid json
	// EncodeJson()
	DecodeJsonData()
}

func EncodeJson()  {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LCO", "abc123", []string{"web", "react"}},
		{"JS Bootcamp", 199, "LCO", "def123", []string{"frontend", "js"}},
		{"MERN Bootcamp", 399, "LCO", "BCD123", nil},
	}

	// package this data as json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%s\n", finalJson)
}

func DecodeJsonData()  {
	// response, err := http.Get("https://dummyjson.com/products/1")

	response := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "LCO",
			"tags": [
					"web",
					"react"
			]
        }
	`)

	// if err != nil {
	// 	panic(err)
	// }

	// defer response.Body.Close()

	// content, err := io.ReadAll(response.Body)

	// if err != nil {
	// 	panic(err)
	// }

	var lcoCourse course
	isValid := json.Valid(response)

	if isValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(response, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("No the data is not valid JSON")
	}
	fmt.Println(string(response))

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(response, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("%v : %v", k, v)
	}
}