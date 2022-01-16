package main

import "fmt"

func main()  {
	// vai executar a func2 primeiro e depois a func1
	defer func1()
	func2()	
}

func func1()  {
	fmt.Println("func1")
}

func func2()  {
	fmt.Println("func2")
}