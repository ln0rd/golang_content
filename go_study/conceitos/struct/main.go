package main

import "fmt"

func main() {
	var u user

	u.age = 26
	u.name = "Gutz"

	fmt.Println(u)

	ad := address{street: "street of fools", number: 10}
	var u2 user = user{"hinata", 16, ad}
	fmt.Println(u2)

	u3 := user{age: 21}
	fmt.Println(u3)
}

// uma colecao de campos, que tem nome e tipo
type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint32
}
