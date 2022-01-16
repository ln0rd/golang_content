package main

import "fmt"

func main()  {
	// meu canal tem capacidade de 2
	canal := make(chan string, 2)
	canal <- "hi"
	canal <- "hi 2"


	msg1 := <-canal
	msg := <-canal
	fmt.Println(msg1)
	fmt.Println(msg)
}