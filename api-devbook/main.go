package main

import (
	"api-devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	fmt.Println("Rodando API")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}