package main

import (
	"fmt"
	"modulo/helper"

	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("Hi")
	helper.Write()
	erro := checkmail.ValidateFormat("leo@leo.com")

	fmt.Println(erro)
}