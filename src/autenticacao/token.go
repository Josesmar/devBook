package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go" // import usando alias
)

// CriarToken retorna um token assinado com as permissões de usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() //Unix vai devolver a quantidade de miliessegundos 15713231112
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes) // Metodo de assinatura SigningMethodHS256 + as permissoes
	return token.SignedString([]byte(config.SecretKey))            // secret
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // Valid Verifica se o token está expirado ou não
		return nil
	}

	return errors.New("Token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 { //Verificando se existe dois valores
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
