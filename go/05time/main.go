package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Time study of golang ")

	currTime := time.Now();
	fmt.Println(currTime)
	fmt.Println(currTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(1999, time.June, 1, 7, 7, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
	
}