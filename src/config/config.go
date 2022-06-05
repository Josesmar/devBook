package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco é a string de conexão com o MySql
	StringConexaoBanco = ""

	//Porta onde a API vai estar rodando
	Porta = 0

	IP = 0
)

// Carrregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil { //gotoenv ler o arquivo
		log.Fatal(erro) //Se tiver algum problema terá que parar (matar a execução mesmo) a api
	}

	/*
		Atoi -> é a mesma coisa que o ParseInt
		os.Getenv -> pega os valores do arquivo .env
	*/
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
