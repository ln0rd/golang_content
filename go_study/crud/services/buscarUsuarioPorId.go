package services

import (
	"crud/banco"
	"crud/errorValidation"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarUsuarioPorId(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	// parse int recebe 3 params, segundo base 10(decimal), terceiro in32, int64 etc
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	errorValidation.ErrorResponseValidation(err, w, "Não foi identificado nenhum ID")

	db, err := banco.Conectar()
	errorValidation.ErrorResponseValidation(err, w, "Erro ao conectar o banco")

	linha, err := db.Query("select * from usuarios where id = ?", ID)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao criar Query")

	var user usuario
	if linha.Next() {
		err := linha.Scan(&user.ID, &user.Nome, &user.Email)
		errorValidation.ErrorResponseValidation(err, w, "Erro ao buscar usuario")

	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao realizar enconder")
}