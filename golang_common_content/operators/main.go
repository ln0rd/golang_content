package main

import (
	"fmt"
)

func main() {
	// ARITIMETCS
	// + - / * %
	fmt.Println(calculate(6, 8))

	//
	var num1 int8 = 2
	var num2 int16 = 3
	// mismatched types int8 and int16
	// WrongSumBecauseTypes = num1 + num2
	fmt.Println(num1, num2)

	// attribution
	var attr1 string = "- -"
	attr2 := 2
	fmt.Println(attr1, attr2)

	// relational operators
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 1)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// logical operators
	fmt.Println("------")
	fmt.Println(1 < 2 && 2 < 3) // and
	fmt.Println(1 > 2 || 3 > 2) // or
	fmt.Println(!true)          // negation

	// unaries operators
	numberTest := 1
	fmt.Println("------")
	numberTest++
	fmt.Println(numberTest)
	numberTest--
	fmt.Println(numberTest)
	numberTest += 10
	fmt.Println(numberTest)

	// ternary operator doesn't exist in Go
	// text := numero > 5 ? "Bigger than five" : "less than five"
	if numberIf > 5 {
		text := "greater than five"
	} else {
		text := "less than five"
	}

}

func calculate(a, b int16) (int16, int16, int16, int16, int16) {
	sum := a + b
	sub := a - b
	div := a / b
	mult := a * b
	mod := a % b

	return sum, sub, div, mult, mod
}
