package packages

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func WritePackage() {
	fmt.Println("Writing from package file")
	checkEmail()
}

func checkEmail() {
	erro := checkmail.ValidateFormat("leo@gmail1.com!")
	fmt.Println(erro)
}