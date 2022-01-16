package main

import "fmt"

func main()  {
	

	func (param string)  {
		fmt.Println("Anonymous function - name:", param)
	}("ln0rd")

	returned := func(text string) string {
		return fmt.Sprintf("String %s", text)	
	}("ln0rd")

	fmt.Println(returned)
}