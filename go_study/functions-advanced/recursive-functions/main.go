package main

import "fmt"

// função recursiva é uma função que chama ela mesma
func main()  {
	
	// sequencia de fibonati
	// proximo numero é sempre a soma dos dois numeros anteriores
	posicao := uint(10)
	fmt.Println( fibonacci(posicao) )
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}