package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array [5]int16
	fmt.Println(array)

	array[0] = 8
	fmt.Println(array)

	arr := [5]string{"position 1", "position 2", "position 3", "position 4", "position 5"}
	arr[0] = "test"
	fmt.Println(arr)

	// the size of this array is going to be the number of passed numbers in constructor
	arr2 := [...]int32{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := arr[1:3]
	fmt.Println(arr3)

	// SLICE
	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice2)
	slice := []int{}
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(slice[0]))

	//INTERNAL ARRAY
	fmt.Println("-------------------")
	// make( type of array, size, max capacity )
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 12.5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacity

}
