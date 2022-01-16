package main

import (
	"errors"
	"fmt"
)

func main() {
	// int16, int32, int64
	// int, depends your architecture machine
	var num int8 = -100
	fmt.Println(num)

	// unsigned int
	// int without signal
	var unum uint8 = 100
	fmt.Println(unum)

	// alias
	// int32 == rune
	var nrune rune = 10
	fmt.Println(nrune)

	// int8 == byte
	var nbyte byte = 8
	fmt.Println(nbyte)

	// floats: float32, float64
	var fnum float32 = 10.5
	var fnum2 float64 = 250.5

	fmt.Println(fnum, fnum2)

	// String
	var str1 string = "Test String"
	fmt.Println(str1)

	// Char doesn't have type in go
	char := 'a'
	fmt.Println(char)

	// booleans
	var boolean bool = true
	var booleanF bool = false

	fmt.Println(boolean, booleanF)

	// type error
	var err error
	fmt.Println(err)

	var er error = errors.New("Internal Error")
	fmt.Println(er)

}
