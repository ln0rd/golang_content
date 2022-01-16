package services

import (
	"crud/banco"
	"crud/errorValidation"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeletarUsuario(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao identificar Id")

	db, err := banco.Conectar()
	errorValidation.ErrorResponseValidation(err, w, "Erro ao conectar no banco")
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	errorValidation.ErrorResponseValidation(err, w, "Erro ao criar statament")
	defer statement.Close()

	_, err = statement.Exec(ID)
	errorValidation.ErrorResponseValidation(err, w, "Erro ao executar statament")

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(fmt.Sprintf("Deleted id %d", ID))
	errorValidation.ErrorResponseValidation(err, w, "Erro ao criar os json para retornar")
}