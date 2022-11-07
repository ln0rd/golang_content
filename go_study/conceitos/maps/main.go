package main

import "fmt"

func main() {

	user := map[string]string{
		"name": "leo",
		"age":  "26",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "leo",
			"last":  "bull",
		},
	}

	fmt.Println(user2)
}
