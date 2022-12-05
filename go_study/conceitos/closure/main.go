package main

import "fmt"

func main() {
	text := "Dentro da funcao main"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()
}

// closure é uma funcao que retorna outra funcao para ser executada
func closure() func() {
	var text string = "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}
