package main

import "fmt"

func main()  {
	// cria um slice um slice 
	slice := make([]float32, 10, 11)
	fmt.Println(slice)
	fmt.Println( "capacity:", cap(slice), "Size:", len(slice))

	slice = append(slice, 15)
	fmt.Println( "capacity:", cap(slice), "Size:", len(slice))

	slice = append(slice, 20)
	fmt.Println( "capacity:", cap(slice), "Size:", len(slice))


	// criando um slice sem o parametro da capacidade
	slice2 := make([]float32, 0)
	fmt.Println("")
	fmt.Println("slice2:", slice2)
	fmt.Println( "capacity:", cap(slice2), "Size:", len(slice2))

	slice2 = append(slice2, 1)
	fmt.Println("slice2:", slice2)
	fmt.Println( "capacity:", cap(slice2), "Size:", len(slice2))

	slice2 = append(slice2, 2)
	slice2 = append(slice2, 3)
	fmt.Println("slice2:", slice2)
	fmt.Println( "capacity:", cap(slice2), "Size:", len(slice2))
}