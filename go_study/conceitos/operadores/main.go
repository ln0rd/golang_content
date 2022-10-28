package main

import "fmt"

func main() {

	// ARITIMETICOS
	// + - / * %

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 1)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// OPERADORES LOGICOS
	fmt.Println(true && false) // todos eles forem true
	fmt.Println(true || false) // se um deles for verdadeiro
	fmt.Println(!true)         // vira false
	fmt.Println(!false)        // vira true

	// OPERADORES UNARIOS
	n := 10
	n++
	n--
	n += 2
	n -= 2
	n *= 2
	n /= 2
	n %= 2
}
