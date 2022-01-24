package helper

import (
	"fmt"
	"golang_study/packages/content"
)

func WriteHelper() {
	fmt.Println("Writing from helper file")
	WriteHelper2()
	content.WriteHelpertwo()
}