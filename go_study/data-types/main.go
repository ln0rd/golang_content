package main

import (
	"errors"
	"fmt"
)

func main()  {
	//numbers
	var num1 int8 	= 100
	var num2 int16 	= 10000
	var num3 int32 	= 1000000000
	var num4 int64 	= 1000000000000000099
	var unu1 uint8 	= 100
	var unu2 uint16 = 10000
	var unu3 uint32 = 1000000000
	var unu4 uint64 = 1000000000000000099
	// uint - Unsigned integer
	fmt.Println(num1, num2, num3, num4)
	fmt.Println(unu1, unu2, unu3, unu4)
	// int32 == rune
	// int8 == byte

	//floats
	var float1 float32 = 10.111111111
	var float2 float64 = 10.111111111
	fmt.Println(float1, float2)

	//strings
	var str1 string = "text"
	str2 := "text"
	fmt.Println(str1, str2)

	//char - ASC table
	char := 'B'
	fmt.Println(char)

	// value 0
	var text string
	fmt.Println(text)
	var num int8
	fmt.Println(num)

	//boolean
	var boolean bool = true
	var boolean2 bool = false
	fmt.Println(boolean, boolean2)

	//error
	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}