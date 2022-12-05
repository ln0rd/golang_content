package main

import "fmt"

func main() {
	var result bool = studentApproved(6, 6)
	fmt.Println(result)

	fmt.Println("pos execucao")
}

func studentApproved(n1, n2 float64) bool {
	defer recoverFromPanic()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// nunca pode ser 6, se for 6 ele entra em panico e para esse processo
	// quando panic e executado ele chama antes todas as funcoes que tem defer
	panic("exatamente 6")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando do panic continuando daqui")
	}
}
