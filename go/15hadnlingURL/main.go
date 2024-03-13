package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.google.com/learn?coursename=reactjs&paymentid=gcfgff"

func main()  {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	res, _ := url.Parse(myUrl)

	fmt.Println(res.Scheme)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)
 
	qparams := res.Query() 
	fmt.Printf("The type of query params are %T\n", qparams)
	fmt.Println(qparams)

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}