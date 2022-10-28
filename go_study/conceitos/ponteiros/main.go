package main

import "fmt"

func main() {

	// valor é atribuido mas não é passado como ponteiro, por isso apenas uma
	// variavel e atualizada
	var vari1 int32 = 10
	var vari2 int32 = vari1

	vari1++
	fmt.Println(vari1, vari2)

	var vari3 int16 = 10
	var ponteiro *int16 = &vari3

	vari3++

	// ele mostra o endereco de memoria
	fmt.Println(vari3, ponteiro)

	// desferenciação
	fmt.Println(vari3, *ponteiro)
}
