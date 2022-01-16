package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cat struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age int32 `json:"age"`
}

func main() {
	creatingJSONCatFromStruct()
	creatingJSONCatFromMap()
}

func creatingJSONCatFromMap() map[string]string  {
	cat := map[string]string {
		"name": "Tobi",
		"race": "Persian",
		"age": "4",
	}

	// trasnforming cat in array of bytes
	cJSON, erro := json.Marshal(cat)
	if erro != nil {
		log.Fatal("error to create json")
	}

	// passando um map para um JSON
	fmt.Println(cJSON)
	fmt.Println(bytes.NewBuffer(cJSON))

	return cat
}

func creatingJSONCatFromStruct() cat {
	cat := cat{ "Charlie", "not defined", 2 }
	fmt.Println(cat)

	// trasnforming cat in array of bytes
	cJSON, erro := json.Marshal(cat)
	if erro != nil {
		log.Fatal("error to create json")
	}

	// passando um struct para um JSON
	fmt.Println(cJSON)
	fmt.Println(bytes.NewBuffer(cJSON))

	return cat
}
