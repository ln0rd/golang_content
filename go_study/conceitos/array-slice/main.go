package main

import (
	"fmt"
	"reflect"
)

func main() {

	// ao criar um array vc determina uma quantia limite
	var arr1 [5]int
	fmt.Println(arr1)
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5

	fmt.Println(arr1)

	arr2 := [3]string{"pos1", "pos2", "pos3"}
	fmt.Println(arr2)

	// o tamanho do array e de acordo com a quantia de valores passados ao criar.
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr3)

	// slice
	slice := []int16{}
	fmt.Println(slice)

	slice = append(slice, 1, 2, 3)
	slice = append(slice, 4)

	fmt.Println(slice)

	// slice e array sao diferentes
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr3))

	// pegar intervalos, do 1 ao 3
	fmt.Println(slice[1:3])

	// array internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)      // vai ter tamanho de 10 e capacidade de 15
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // tamanho
	fmt.Println(cap(slice4)) // capacidade

	slice4 = append(slice4, 10, 11, 12)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // tamanho
	fmt.Println(cap(slice4)) // capacidade
}
