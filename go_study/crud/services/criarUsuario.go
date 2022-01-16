package services

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type usuario struct {
	ID uint32 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request)  {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user1 usuario

	err = json.Unmarshal(payload, &user1)
	if err != nil {
		w.Write([]byte("Erro ao converter usuário para o struct"))
		return 
	}

	fmt.Println(user1)

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return 
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar Statement"))
		return 
	}
	defer statement.Close()

	insercao, err := statement.Exec(user1.Nome, user1.Email)
	if err != nil {
		w.Write([]byte("Erro ao inserir no banco de dados"))
		return 
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao buscar o id inserido"))
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso %d", idInserido)))
}