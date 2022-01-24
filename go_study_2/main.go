package main

import (
	"fmt"
	"go_study_2/packages"
	"go_study_2/packages/helper"
)

func main() {
	fmt.Println("Writing from main file")
	packages.WritePackage()
	helper.WriteHelper()
}