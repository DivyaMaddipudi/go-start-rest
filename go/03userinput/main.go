package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	username := "Divya"
	fmt.Println(username)

	//If i know type use Scan
	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	fmt.Println("Nice talking to you, ", name)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our pizza:")

	// Comma ok syntax || error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T ", input)




}