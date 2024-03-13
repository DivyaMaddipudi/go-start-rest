package main

import "fmt"

func main()  {
	fmt.Println("If else in go lang")
	loginCount := 11

	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if(loginCount > 10 && loginCount < 15) {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}