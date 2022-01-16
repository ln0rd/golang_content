package main

import (
	"fmt"
)

type user struct {
	name string
	age uint8
}

func (u user) save() {
	fmt.Printf("Saving user %s \n", u.name)
}

func (u *user) birthday() uint8 {
	u.age++
	return u.age
}

func (u user) verifyAge() bool  {
	return u.age > 18
}

func main()  {
	user1 := user{ name: "leo", age: 25 }
	fmt.Println(user1)
	user1.save()

	if user1.verifyAge() {
		fmt.Println("This user is over 18 years old")
	}

	user1.birthday()
	fmt.Println(user1.age)
}