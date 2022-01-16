package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cat struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age int32 `json:"age"`
	Size int32 `json:"-"` // it will be ignored
}

func main() {
	creatingStructCatFromJSON()
}

func creatingStructCatFromJSON()  {
	catFromJson := []byte(`{ "name":"Charlie", "race":"not defined", "age":2 }`)

	var cat cat

	err := json.Unmarshal(catFromJson, &cat)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cat)
}

func creatingStructCatFromMap()  {
	catFromJson := []byte(`{ "name":"Charlie", "race":"not defined", "age":2 }`)

	cat := make(map[string]string)

	err := json.Unmarshal(catFromJson, &cat)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cat)
}
