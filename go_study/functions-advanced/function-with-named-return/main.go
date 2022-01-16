package main

import "fmt"

func main()  {
	soma, sub := calculo(10, 5)
	fmt.Println("Soma:", soma, "Sub:", sub)
}

// retorno nomeado
func calculo(n1 int, n2 int) (soma int, sub int)  {
	soma = n1 + n2
	sub = n1 - n2
	return
}