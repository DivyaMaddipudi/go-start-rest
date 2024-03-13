package main

import "fmt"

func main()  {
	fmt.Println("Functions")
	greet()
	result, message := proAdder(3, 4, 5, 6)
	fmt.Println("result is: ", result)
	fmt.Println("message is: ", message)

}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}
	return total, "second value from function"
}

func add(one int, two int) int {
	return one + two
}

func greet2()  {
	fmt.Println("Another method")
}

func greet() {
	fmt.Println("Namastey from golang")
}