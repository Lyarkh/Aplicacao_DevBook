package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin irá renderizar a tela de Login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroUsuario vai carregar a pagina de usuario
func CarregarPaginaDeCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}