package main

import "fmt"

func main() {
	user1 := user{"user 1", 6}
	user1.save()

	var isLegalAge bool = user1.isLegalAge()
	fmt.Println(isLegalAge)
}

type user struct {
	name string
	age  int8
}

// metodo
func (u user) save() {
	fmt.Printf("Salvando %s", u.name)
}

func (u user) isLegalAge() bool {
	if u.age > 18 {
		return true
	} else {
		return false
	}
}
