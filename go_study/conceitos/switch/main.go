package main

import "fmt"

func main() {
	var result string = diaDaSemana(2)
	fmt.Println(result)
}

func diaDaSemana(num int32) string {
	switch num {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terca"
	default:
		return "invalid"
	}
}
