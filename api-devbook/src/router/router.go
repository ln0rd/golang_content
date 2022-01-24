package router

import "github.com/gorilla/mux"

// *mux.Router é o retorno
func Gerar() *mux.Router {
	return mux.NewRouter()
}