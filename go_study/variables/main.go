package main

import "fmt"

func main()  {
	var variable1 string = "leo"
	fmt.Println(variable1)

	variable2 := "vlad"
	fmt.Println(variable2)

	var1, var2, var3 := "var1", "var2", "var3"
	fmt.Println(var1, var2, var3)

	var (
		vr1 string = "sword"
		vr2 string = "short bow"
	)
	fmt.Println(vr1, vr2)

	// invert values
	var1, var2 = var2, var1 
	fmt.Println(var1, var2)
}