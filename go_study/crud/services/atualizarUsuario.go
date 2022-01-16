package services

import (
	"crud/banco"
	"crud/errorValidation"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AtualizarUsuario(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao ler parametro Id")

	bodyRequest, err := ioutil.ReadAll(r.Body)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao ler corpo da requisicao")

	var user usuario
	err = json.Unmarshal(bodyRequest, &user)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao traduzer user do bodyRequest")

	db, err := banco.Conectar()
	errorValidation.ErrorResponseValidation(err, w, "Erro ao conectar ao banco")

	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	errorValidation.ErrorResponseValidation(err, w, "Erro ao criar statement")

	defer statement.Close()

	_, err = statement.Exec(user.Nome, user.Email, ID)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao executar statament")

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode("OK")
	errorValidation.ErrorResponseValidation(err, w, "Erro ao criar os json para retornar")
}