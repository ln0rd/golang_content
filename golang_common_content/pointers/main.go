package main

import "fmt"

func main() {

	// without reference
	var v1 int8 = 10
	var v2 int8 = v1
	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	// with reference
	var v3 int8 = 2
	var pointer *int8 = &v3
	fmt.Println(v3, *pointer)

	v3++
	fmt.Println(v3, *pointer)
}
