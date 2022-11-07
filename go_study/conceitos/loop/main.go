package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// 	i++
	// }

	colors := [3]string{"purple", "green", "yellow"}

	for indice, item := range colors {
		fmt.Println(indice, item)
	}

	for _, item := range colors {
		fmt.Println(item)
	}

	// range no string
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra, string(letra))
	}

	user := map[string]string{
		"name": "leo",
		"age":  "26",
	}

	for chave, valor := range user {
		fmt.Println(chave, valor)
	}
}
