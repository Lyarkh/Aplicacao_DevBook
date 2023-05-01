package controllers

import "net/http"


// CriarUsuario irá criar um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando Usuário!"))
}

// BuscarUsuarios irá buscar todos os usuarios no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Todos os usuários!"))
}

// BuscaUsuario irá buscar o usuário pelo ID
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuario!"))
}

// AtualizarUsuarios irá atualizar um usuarios no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeletarUsuario irá deletar um usuario do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
