package main

import "fmt"
const LogInToken string = "abcdef" // public variable because first character of name is Capital

func main() {
	var username string = "Divya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)


	var smallFloat float64 = 255.45555367223232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// deafult values and some aliases
	var anothervariable string
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type: %T \n", anothervariable)

	// implict
	var website = "google.com"
	fmt.Println(website)

	// no var style 
	numberOfUsers := 30000.0
	fmt.Println(numberOfUsers)

	fmt.Println(LogInToken)
	fmt.Printf("Variable is of type: %T \n", LogInToken)


	
}