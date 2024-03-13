package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate pizza b/w 1 and 5")
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')
	fmt.Println(inp)
	// fmt.Printf("Type of inp %T ", inp)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(inp), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", (numRating + 1))
	}
}