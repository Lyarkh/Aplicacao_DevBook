package router

import "github.com/gorilla/mux"

// Gerar retorna um router com todas as rotas criadas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
