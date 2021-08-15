package middleware

import (
	"cookbook/src/authentication"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar() gin.HandlerFunc {
	return func(c *gin.Context) {
		if erro := authentication.ValidarToken(c); erro != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao realizar autenticação: %s", erro.Error())})
			c.Abort()
		}
		c.Next()
	}
}
