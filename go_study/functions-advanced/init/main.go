package main

import "fmt"

var n int 

func main()  {
	fmt.Println("main function")	
}

func init()  {
	n = 1
	fmt.Println(`init function`)
	fmt.Println(n)
}
