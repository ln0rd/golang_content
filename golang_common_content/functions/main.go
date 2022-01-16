package main

import (
	"fmt"
)

func main() {
	// simple function
	sum1 := sum(2, 2)
	fmt.Println(sum1)

	// variable with function type
	var sumFunction = sum
	fmt.Println(sumFunction(1, 1))

	// creating a function inside variable
	var funSum = func(a int16, b int16) int16 {
		return a + b
	}
	fmt.Println(funSum(4, 4))

	// two returns
	sum, sub := calculate(10, 8)
	fmt.Println(sum, sub)

	// ignoring an return
	_, sub2 := calculate(4, 9)
	fmt.Println(sub2)
}

func sum(a int8, b int8) int8 {
	return a + b
}

func calculate(a, b int16) (int16, int16) {
	sum := a + b
	sub := a - b
	return sum, sub
}
