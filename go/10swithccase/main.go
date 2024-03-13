package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	
	// rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	diceNumber := rng.Intn(10) + 1
	fmt.Println("value ", diceNumber)

	switch(diceNumber) {
		case 1: 
			fmt.Println("case 1", diceNumber)
			fallthrough
		case 2: 
			fmt.Println("case 2", diceNumber)
			fallthrough
		case 3: 
			fmt.Println("case 3", diceNumber)
		default:
			fmt.Println("default ", diceNumber)
	}
}