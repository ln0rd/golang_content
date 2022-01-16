package main

import "fmt"

func main()  {
	soma(2,4,6,8,10)
}

func soma(numbers ...int) int {
	// numbers vem como um slice
	current := 0
	for _, number := range numbers {
		fmt.Println("current:", current, "+", number)
		current += number
		fmt.Println("current:", current)
	}

	return current
}