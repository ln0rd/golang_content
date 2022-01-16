package main

import "fmt"

func main()  {
	soma := 1 + 2
	sub := 1 - 2
	divisao := 10 / 2
	mult := 2 * 5
	restoMod := 10 % 2

	fmt.Println(soma, sub, divisao, mult, restoMod)

	var num1 int16 = 10
	var num2 int16 = 25

	soma1 := num1 + num2
	fmt.Println(soma1)


	// attr
	var var1 string = "ln0rd"
	var2 := "ln0rd"
	fmt.Println(var1, var2)

	// relacionais
	// == <= >= ! !=
	fmt.Println(1 > 1)
	fmt.Println(1 < 2)
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 <= 1)

	// logical operators
	// end &&
	// or ||
	// ! invert
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)


	// Unary
	var n int16 = 10
	n++
	n += 15
	n--
	n -= 10
	fmt.Println(n)

	// ternario
	// não há
}