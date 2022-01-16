package main

import "fmt"

func main()  {
	var myuser user
	myuser.age = 25
	myuser.name = "ln0rd"

	myuser2 := user{"ln0rd", 25}
	fmt.Println(myuser2)

	myuser3 := user{age: 25}
	fmt.Println(myuser3)

	andress1 := andress{"Wall Street", 2}
	myuserWithAndress1 := userWithAndress{"ln0rd", 25, andress1}
	fmt.Println(myuserWithAndress1)
}

type user struct {
	name string
	age int8
}

type userWithAndress struct {
	name string
	age int8
	andress andress
}

type andress struct {
	street string
	number int8
}