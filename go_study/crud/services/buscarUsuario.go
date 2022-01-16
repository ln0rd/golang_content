package services

import (
	"crud/banco"
	"crud/errorValidation"
	"encoding/json"
	"net/http"
)

func BuscarUsuarios(w http.ResponseWriter, r *http.Request)  {
	db, err := banco.Conectar()
	errorValidation.ErrorResponseValidation(err,w, "Erro ao conectar ao bancos")
	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	errorValidation.ErrorResponseValidation(err, w, "Erro ao buscar os usuarios")
	defer linhas.Close()

	var users []usuario
	for linhas.Next() {
		var user usuario

		err := linhas.Scan(&user.ID, &user.Nome, &user.Email)
		errorValidation.ErrorResponseValidation(err, w, "Erro enquanto retornava os usuarios da base")

		users = append(users, user)
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(users)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao criar os json para retornar")

	
}