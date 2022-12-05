package main

import "fmt"

func main() {
	var numero int64 = 20
	var invertido int64 = inverterSinal(numero)

	fmt.Println(numero)
	fmt.Println(invertido)

	var novoNumero int = 40
	invertendoPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}

func inverterSinal(numero int64) int64 {
	return numero * -1
}

func invertendoPonteiro(numero *int) {
	*numero = *numero * -1
}
