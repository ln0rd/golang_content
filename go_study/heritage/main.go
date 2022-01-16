package main

import "fmt"

func main()  {
	people1 := people{"Ln0rd", "Nord", 25, 178}
	student1 := student{people1, "leo"}
	fmt.Println(student1.name)
}

type people struct{
	name string
	lastName string
	age int8
	size int16
}

type student struct{
	people
	course string
}