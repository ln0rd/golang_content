package main

import "fmt"

func main()  {
	texto := "inside main function"
	fmt.Println(texto)
	
	newFunction := closure()

	newFunction()
}

func closure() func() {
	texto := "inside closure function"
	
	funcao := func ()  {
		fmt.Println(texto)
	}

	return funcao
}