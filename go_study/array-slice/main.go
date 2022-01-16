package main

import (
	"fmt"
	"reflect"
)

func main()  {
	// array
	var array1[5] int
	array1[0] = 1
	fmt.Println(array1)

	// slice
	slice := []int{}
	fmt.Println(slice)
 	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18)
	fmt.Println(slice)

	// pegar intervalo
	array2 := [5]string{"p1", "p2", "p3", "p4", "p5"} 
	slice2 := array2[1:3]
	fmt.Println(slice2)
}