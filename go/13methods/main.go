package main

import "fmt"

func main()  {
	fmt.Println("Methods in struct")
	divya := User {"Divya", "divya@gmail.com", 24, true}
	fmt.Println(divya)
	divya.GetStatus()
	divya.NewMail()
	fmt.Println(divya)
}

type User struct {
	Name string
	Email string
	Age int
	Status bool
}

func (u User)  GetStatus() {
	fmt.Println("Is user actiove: ", u.Status)
}

func ( u User)  NewMail() {
	u.Email = "satya@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}