package main

import (
	"fmt"
	"golang_study/packages"
	"golang_study/packages/helper"
	"golang_study/variables"
)

func main() {
	// packages
	fmt.Println("Writing from main file")
	packages.WritePackage()
	helper.WriteHelper()
	// variables
	variables.WritingVariable()
}