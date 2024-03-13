package main

import "fmt"

func main()  {
	// var highScores [3]int
	var highScores = []int{234, 800, 100, 300, 500, 1000}
	// highScores := make([]int, 2)
	highScores = append(highScores, 700, 320)
	// highScores[2] = 234	

	fmt.Println(highScores)

	// remove value from slice

	var ind int = 2
	
	highScores = append(highScores[:ind], highScores[ind + 1: ]...)
	fmt.Println(highScores)
	
}