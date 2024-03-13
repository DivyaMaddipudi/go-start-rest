package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/"

func main()  {
	fmt.Println("Web requests")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response is %T\n", response)

	defer response.Body.Close() // callers responsibilty to close this

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
	panic(err)
	}

	content := string(dataBytes)
	fmt.Println(content)

}