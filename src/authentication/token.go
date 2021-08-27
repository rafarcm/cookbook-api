package authentication

import (
	"cookbook/src/config"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CriarToken -> retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64, empresaID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	permissoes["empresaId"] = empresaID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}

// ValidarToken : verifica se o token passado na requisição é valido
func ValidarToken(c *gin.Context) error {
	tokenString := extrairToken(c)
	token, erro := jwt.Parse(tokenString, retornarChaveDeverificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// ExtrairIDs -> retorna o usuarioID e empresaID que esta salvo no token
func ExtrairIDs(c *gin.Context) (uint64, uint64, error) {
	tokenString := extrairToken(c)
	token, erro := jwt.Parse(tokenString, retornarChaveDeverificacao)
	if erro != nil {
		return 0, 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, 0, erro
		}

		empresaID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["empresaId"]), 10, 64)
		if erro != nil {
			return 0, 0, erro
		}

		return usuarioID, empresaID, nil
	}

	return 0, 0, errors.New("token inválido")
}

func extrairToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeverificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
