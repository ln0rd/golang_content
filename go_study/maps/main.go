package main

import "fmt"

func main()  {
	
	user := map[string]string {
		"name": "ln0rd",
		"lastname": "nord",
	}
	fmt.Println(user)
	fmt.Println(user["name"])
	fmt.Println(user["lastname"])

	user2 := map[string]map[string]string {
		"user": user,
		"user1": {
			"name": "leo",
			"lastname": "kruggler",
		},
	}
	fmt.Println(user2)
	fmt.Println(user2["user"]["name"])

	// delete
	delete(user2, "user1")
	fmt.Println(user2)
	
	//add inside map
	user2["user3"] = map[string]string{
		"name": "Adam",
		"lastname": "Sandler",
	}
	fmt.Println(user2)
}