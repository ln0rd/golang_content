package main

import "fmt"

func main()  {
	
	num := 10
	limit := 15

	// if
	if num > limit {
		fmt.Println("bigger than 15")
	} else {
		fmt.Println("less than 15")
	}

	// if init
	if maybeVar := true; maybeVar == true {
		fmt.Println("mayberVar has been atribuited with true and it is True")
	}
}