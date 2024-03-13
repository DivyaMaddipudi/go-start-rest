package main

import "fmt"

func main()  {
	fmt.Println("Pointers")

	// var ptr *string;
	// fmt.Println(ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println(ptr)
	fmt.Println(*ptr, " : ", myNumber*3)

	*ptr = *ptr * 2
	fmt.Println("my Number ", myNumber)

}