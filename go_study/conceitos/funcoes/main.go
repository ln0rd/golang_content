package main

import "fmt"

func main() {
	var soma int16 = somar(4, 4)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("teste")

	result1, result2 := calculos(10, 4)
	fmt.Println(result1, result2)

	_, resultv2 := calculos(10, 4)
	fmt.Println(resultv2)
}

func somar(n1 int16, n2 int16) int16 {
	return n1 + n2
}

func calculos(n1, n2 int32) (int32, int32) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
