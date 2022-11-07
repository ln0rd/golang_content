package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	// cria e ja valida
	if outronumero := numero; outronumero > 0 {
		fmt.Println("maior que 0")
	} else {
		fmt.Println("menor que 15")
	}
}
