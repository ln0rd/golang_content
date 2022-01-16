package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")

	user := map[string]string{
		"name":     "rick",
		"lastanem": "sanches",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"user1": {
			"name":     "jhon",
			"lastname": "constantine",
		},
	}
	fmt.Println(user2)
	fmt.Println(user2["user1"]["lastname"])

	delete(user2, "name")
	delete(user2["user1"], "name")
	fmt.Println(user2)

	user2["user2"] = map[string]string{
		"name":     "gerald",
		"lastname": "of the rivia",
	}

	fmt.Println(user2)
}
