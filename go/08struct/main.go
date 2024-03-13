package main

import "fmt"

func main() {
	fmt.Println("Strcuts in Golang")

	divya := User {"Divya", "divya@gmail.com", true, 24}

	fmt.Printf("Divya details are: %+v\n", divya)

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}