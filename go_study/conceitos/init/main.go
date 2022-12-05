package main

import "fmt"

// antes da função main, pode ser um init por arquivo e nao por pacote
func init() {
	fmt.Println("vem antes da main")

}

func main() {
	fmt.Println(" main")
}
