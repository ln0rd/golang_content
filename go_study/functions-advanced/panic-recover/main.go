package main

import "fmt"

func main()  {
	aproved(6, 6)
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recuperada com sucesso")
	}

}


func aproved(n1, n2 float32) bool  {
	defer recoverFromPanic()
	media := (n1 + n2) / 2

	if media > 6 {
		fmt.Println("approved")
		return true
	} else if media < 6 {
		fmt.Println("Not approved")
		return false
	}

	// panic para tudo e chama os defer antes de parar a aplicação
	panic("PANIC, it's 6")
}