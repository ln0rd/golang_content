package main

import "fmt"

func main()  {
	
	var v1 int = 10
	var v2 int = 11
	var v3 *int

	// atribuicao
	v1 = v2
	fmt.Println("V1:", v1," V2:" ,v2, "V3:", v3)

	//ponteiro
	v3 = &v2
	fmt.Println(v3)
	fmt.Println(*v3)
}