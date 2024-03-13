package main

import "fmt"

func main()  {
	days := []string{}
	days = append(days, "Monday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")
	fmt.Println(days)

	// for i :=0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i], i)
	// }

	for ind, day := range days{
		if(ind == 2) {
			goto divya
		}
		fmt.Printf("%v %v\n", ind, day)
	}

	divya:
		fmt.Println("User divya is called")
}