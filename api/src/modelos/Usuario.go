package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuario do banco de dados
type Usuario struct {
	ID       uint64    `json:"id, omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar faz a validação dos campos e formatação do struct usuario para enviar para o DB
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome deve ser preenchido")
	}
	if usuario.Nick == "" {
		return errors.New("o nick deve ser preenchido")
	}
	if usuario.Email == "" {
		return errors.New("o email deve ser preenchido")
	}
	if usuario.Senha == "" {
		return errors.New("a senha deve ser preenchido")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
