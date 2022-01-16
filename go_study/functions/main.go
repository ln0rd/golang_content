package main

import "fmt"

func main()  {
	soma := somar(10,20)
	fmt.Println(soma)

	//func inside var
	var f = func (t string)  {
		fmt.Println(t)
	}
	f("I'm Walking")

	//getting two returns
	result1, result2 := calculate(1,2)
	fmt.Println(result1, result2)

	//just want result2
	_, result := calculate(1,2)
	fmt.Println(result)
}

func somar(number1 int8, number2 int8) int8 {
	return number1 + number2
}

func calculate(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}