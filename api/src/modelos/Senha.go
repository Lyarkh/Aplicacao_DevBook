package modelos

// Senha representa o formato de requisição de alteração de senha
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}