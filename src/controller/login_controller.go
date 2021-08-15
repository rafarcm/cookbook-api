package controller

import (
	"cookbook/src/authentication"
	"cookbook/src/model"
	"cookbook/src/security"
	"cookbook/src/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type loginController struct {
	usuarioService service.UsuarioService
}

// LoginController : representa o contrato de LoginController
type LoginController interface {
	Login(*gin.Context)
}

//NewLoginController -> retorna um LoginController
func NewLoginController(u service.UsuarioService) LoginController {
	return loginController{
		usuarioService: u,
	}
}

// Login: realiza o login de um usuário
func (l loginController) Login(c *gin.Context) {
	log.Print("[LoginController]...Realizando Login")

	var login model.Login
	if erro := c.ShouldBindJSON(&login); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	usuario, erro := l.usuarioService.FindByUsername(login.Username)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao realizar login: %s", erro.Error())})
		return
	}

	if usuario.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username inválido"})
		return
	}

	if erro = security.VerificarSenha(usuario.Senha, login.Senha); erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha inválida"})
		return
	}

	token, erro := authentication.CriarToken(usuario.ID, usuario.EmpresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao realizar login: %s", erro.Error())})
		return
	}

	usuarioID := strconv.FormatUint(usuario.ID, 10)

	c.JSON(http.StatusOK, gin.H{"data": model.DadosAutenticacao{ID: usuarioID, Token: token}})
}
