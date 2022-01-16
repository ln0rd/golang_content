package main

import "fmt"

func main() {

	var person person = person{name: "Rick", lastname: "Sanches", age: 58, size: 170}
	fmt.Println(person)

	var std student = student{person: person, course: "Scientist", university: "Intergalactic university"}
	fmt.Println(std)

}

type person struct {
	name     string
	lastname string
	age      int16
	size     int32
}

// HERITAGE
type student struct {
	person
	course     string
	university string
}
