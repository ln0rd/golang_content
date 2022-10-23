package main

import "fmt"

func main() {
	var name string = "onoda"
	var bike string = "bmc"
	fmt.Println(name, bike)

	var (
		age   int32  = 27
		phone string = "999999999"
	)
	fmt.Println(age, phone)

	const size_cm int32 = 175

	//Invertendo valores entre duas variaveis
	name, bike = bike, name
}
