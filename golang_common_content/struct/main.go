package main

import "fmt"

func main() {

	// struct is like class in another language
	var us user
	fmt.Println(us)
	us.name = "Guts"
	us.age = 37
	fmt.Println(us)

	us2 := user{"Misaki", 19}
	fmt.Println(us2)

	var us3 user = user{"Madara", 40}
	fmt.Println(us3)

	var us4 user = user{age: 40}
	fmt.Println(us4)

	var weapon1 weapon = weapon{kind: "lance", size: "170cm", user: us2}
	fmt.Println(weapon1)

	var weapon2 weapon = weapon{kind: "sword", size: "90cm", user: user{"Guts", 37}}
	fmt.Println(weapon2)
}

type user struct {
	name string
	age  int8
}

type weapon struct {
	kind string
	size string
	user user
}
