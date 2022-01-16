package main

import (
	"fmt"
)

func main() {
	var var1 string = "var 1"
	fmt.Println(var1)

	var2 := "var 2"
	fmt.Println(var2)

	var (
		name     string = "Sherlock"
		lastname string = "Holmes"
	)

	fmt.Println(name, lastname)

	person1, person2 := "Sam porter Bridges", "Death Strading"

	fmt.Println(person1, person2)
}
