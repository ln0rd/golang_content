package main

import (
	"errors"
	"fmt"
)

func main() {
	// numeros inteiros
	// int8, int16, int32, int64
	// int sem especificar o valor ele usa a arquitetura do computador como base.
	// uint unsygned integer, inteiro sem sinal, nao pode ser negativo somente positivo
	// rune e um alias para int32
	var nu1 int8 = 1
	fmt.Println(nu1)

	var nu2 int64 = 1
	fmt.Println(nu2)

	var nurune rune = 1
	fmt.Println(nurune)

	var nubyte byte = 8
	fmt.Println(nubyte)

	// numeros reais
	// float32, float64

	var nu3 float32 = 12.345
	fmt.Println(nu3)

	var nu4 float64 = 13.90909
	fmt.Println(nu4)

	// texto
	// string, char

	var c string = "a"
	fmt.Println(c)

	// valor ZERO, valor atribuido quando voce nao inicializa ela
	var texto string
	fmt.Println(texto)

	var inteirovazio int32
	fmt.Println(inteirovazio)

	var err error
	fmt.Println(err)

	err = errors.New("erro interno")
	fmt.Println(err)
}
