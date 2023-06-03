package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

// CarregarTelaDeLogin irá renderizar a tela de Login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroUsuario vai carregar a pagina de usuário
func CarregarPaginaDeCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal carrega a pagina principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeError(w, response)
		return
	}
	utils.ExecutarTemplate(w, "home.html", nil)
}
