package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Starting email validator")

	err := checkmail.ValidateFormat("test@gmail.com")
	fmt.Println(err)
}
