package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	// PerformGetRequest("https://dummyjson.com/products/1");
	// PerformPostJsonRequest(); //- we will get error because i'm not using the post URL
	// PerformPostFormRequest(); //- we will get error because i'm not using the post URL
}

func PerformGetRequest(myurl string)  {

	response, error := http.Get(myurl)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder

	bytesData, error := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(bytesData)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println("Response string is: ", responseString.String())
	
	if error != nil {
		panic(error)
	}
	fmt.Println(string(bytesData))
}

func PerformPostJsonRequest() {
	const myurl = ""

	requestBody := strings.NewReader(`{
		"coursename" : "Lets go with golang",
		"pirce": 0,
		"platform": "LCO"
	}`)

	fmt.Println("Sending request: ", requestBody)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func PerformPostFormRequest()  {

	const myurl = ""

	// form data
	data := url.Values{}
	data.Add("firstname", "Satyadivya")	
	data.Add("lastname", "Maddipudi")	
	data.Add("email", "divya@gmail.com")	

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("response we sent is: ", response)

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
