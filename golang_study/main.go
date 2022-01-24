package main

import (
	"fmt"
	"golang_study/packages"
	"golang_study/packages/helper"
)

func main() {
	fmt.Println("Writing from main file")
	packages.WritePackage()
	helper.WriteHelper()
}