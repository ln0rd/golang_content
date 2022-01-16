package main

import (
	"fmt"
	"time"
)

func main()  {
	
	i := 0

	for i < 10 {
		// durma por um segundo
		i++
		time.Sleep(time.Second)
		fmt.Println("Incrementando:", i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando:", j)
	}

	//example with array
	names := [3]string{"ln0rd", "keanu", "spider"}

	// for array
	for indice, name := range names {
		fmt.Println("indice:", indice, "name:", name)
	}

	// for string
	for _, letra := range "WORD" {
		fmt.Println( string(letra))
	}

	//for map
	weapon := map[string]string{
		"name": "sword",
		"damage": "10",
	}

	for indice, item := range weapon {
		fmt.Println("indice:", indice, "damage:", item)
	}

}