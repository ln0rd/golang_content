package main

import "fmt"

func main()  {
	num := 20
	invertSignalUsingPointerToOriginalVar(&num)
	fmt.Println(num)
}

func invertSignal(num int) int {
	return num * -1
}

func invertSignalUsingPointerToOriginalVar(num *int) {
	*num = *num * -1
}

// & I can say to method it is a memory andress
// * it does internal scope understand this is pointer