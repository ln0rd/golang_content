package main

import "fmt"

func main() {
	_, sub := calculos(10, 2)
	fmt.Println(sub)

	sum, sub := calculos(10, 2)
	fmt.Println(sum, sub)

	var total int32 = variativas(12, 4, 20, 15)
	fmt.Println(total)

	// funcao anonima
	func(text string) string {
		return fmt.Sprintf("Recebido -> %s %d", text, 2)
	}("parametro")
}

// retorno nomeado com multiplo retorno
func calculos(n1, n2 int) (soma int, sub int) {
	return n1 + n2, n1 - n2
}

// recebendo inumeros parametros
func variativas(num ...int32) int32 {
	var total int32 = 0
	for _, item := range num {
		total += item
	}

	return total
}
