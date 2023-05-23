package controllers

import "net/http"

// CarregarTelaDeLogin irá renderizar a tela de Login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login"))
}
