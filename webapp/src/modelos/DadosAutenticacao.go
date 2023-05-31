package modelos

// DadosAutenticacao contem o ID e token do usuario autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
