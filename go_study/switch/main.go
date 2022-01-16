package main

import "fmt"

func main()  {
	var dayOfWeek string = DayOfWeek(1)	
	fmt.Println(dayOfWeek)
}

// não tem break

func DayOfWeek(day int8) string {
	switch day {
	case 1:
		fmt.Println("Domingo")
		//caso entre no caso 1, vai acionar o caso 2
		// nao pode ter retorno
		fallthrough
	case 2: 
		return "Segunda"
	case 3:
		return "Terca"
	case 4: 
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7: 
		return "Sabado"
	default:
		return "Day undefined"
	}
}