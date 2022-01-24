package variables

import "fmt"

func WritingVariable() {
	// applying type
	var name string = "varivable one"
	fmt.Println(name)

	// occulting type inference
	lastname := "variable two"
	fmt.Println(lastname)

	// declaring more than one variable
	var (
		firstname string = "mars"
		secondname string = "venus"
	)
	fmt.Println(firstname, secondname)

	// declaring more than one varibale by inference
	planet, anotherplanet := "jupiter", "Uranus"
	fmt.Println(planet, anotherplanet)

	// changing variables
	firstname, secondname = secondname, firstname
	fmt.Println(firstname, secondname)
}