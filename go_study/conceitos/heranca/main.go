package main

import "fmt"

func main() {
	var p people = people{name: "Sheldon", age: 41, size: 186}
	var s student = student{p, "ads", "unifran"} // heranca

	fmt.Println(s.name)
	fmt.Println(s.people)
	fmt.Println(s.university)
	fmt.Println(s.age)

	fmt.Println(s)
}

type people struct {
	name string
	age  uint32
	size uint32
}

type student struct {
	people
	course     string
	university string
}
