package main

import "fmt"

func main() {
	// fibonacci
	// 1 1 2 3 5 8 13

	var posicao uint32 = fibonacci(10)
	fmt.Println(posicao)
}

func fibonacci(posicao int32) uint32 {
	if posicao <= 1 {
		return uint32(posicao)
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
