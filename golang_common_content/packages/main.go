package main

// importing external and internal packages
import (
	"fmt"
	"helper"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("From main")

	helper.Write()

	erro := checkmail.ValidateFormat("lslslsls@gmail.com")

	fmt.Println(erro)
}
