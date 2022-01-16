package main

import (
	"crud/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	
	router := mux.NewRouter()
	router.HandleFunc("/users", services.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/users", services.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", services.BuscarUsuarioPorId).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", services.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", services.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Running at port 5000")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal(err)
	}
}