package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizandoa rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"` //omit empty retira o que não tiver valor
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempt"`
	CriadoEm time.Time `json:"criadaEm,omitempty"`
}

//reparar vai chamr os métodos para validar e fomatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {	
	if usuario.Nome == "" {
		return errors.New("O nme é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nik é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pde estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}
	
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é origatória e não pode estar em branco")
	}

	return nil //valor zero do erro
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
