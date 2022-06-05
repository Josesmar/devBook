package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuariosF
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

/*
Criar insere um usuário no banco de dados
o retorno é uint64 pois após ser criado deverá ser retornado um id
*/
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ? ,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario) // Primeiro o slice que vamos adicionar e depois o item que vai ser adicionado
	}

	return usuarios, nil // Retornando o slice de usuários sem erro
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorID(usuarioID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?", usuarioID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro // Retornando um usuário que tem todos os seus valores zero
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera as informações do usuário no banco de dados
func (repositorio Usuarios) Atualizar(usuarioID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ? ",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID); erro != nil {
		return erro
	}
	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Usuarios) Deletar(usuarioID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioID); erro != nil {
		return erro
	}
	return nil
}

// BuscarUsuarioPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarUsuarioPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}
