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
	"gorm.io/gorm"
)

type usuarioController struct {
	usuarioService service.UsuarioService
}

// UsuarioController : representa o contrato de UsuarioController
type UsuarioController interface {
	AddUsuario(*gin.Context)
	UpdateUsuario(*gin.Context)
	DeleteUsuario(*gin.Context)
	FindUsuarioById(*gin.Context)
	GetAllUsuarios(c *gin.Context)
}

//NewUsuarioController -> retorna um UsuarioController
func NewUsuarioController(r service.UsuarioService) UsuarioController {
	return usuarioController{
		usuarioService: r,
	}
}

// AddUsuario: adiciona uma nova Usuario
func (u usuarioController) AddUsuario(c *gin.Context) {
	log.Print("[UsuarioController]...Adicionando usuário")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	var usuario model.Usuario
	if erro := c.ShouldBindJSON(&usuario); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	senhaComHash, erro := security.Hash(usuario.Senha)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar usuário: %s", erro.Error())})
		return
	}
	usuario.Senha = string(senhaComHash)

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao adicionar usuário: %s", erro.Error())})
		return
	}
	if empresaID != usuario.EmpresaID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Não é possível adicionar um usuário para uma empresa que não a sua"})
		return
	}

	usuario, erro = u.usuarioService.WithTrx(txHandle).Save(usuario)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar usuário: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": usuario})
}

// UpdateUsuario : atualiza a Usuario pelo seu id
func (u usuarioController) UpdateUsuario(c *gin.Context) {
	log.Print("[UsuarioController]...Atualizando usuário")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	usuarioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar usuário: %s", erro.Error())})
		return
	}

	var usuario model.Usuario
	if erro := c.ShouldBindJSON(&usuario); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar usuário: %s", erro.Error())})
		return
	}

	usuarioTokenID, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao atualizar usuário: %s", erro.Error())})
		return
	}
	if usuarioID != usuarioTokenID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Não é possível atualizar um usuário que não o seu"})
		return
	}
	if empresaID != usuario.EmpresaID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Não é possível atualizar um usuário para uma empresa que não a sua"})
		return
	}

	usuario.ID = usuarioID
	usuario, erro = u.usuarioService.WithTrx(txHandle).Update(usuario)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao atualizar usuário: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usuario})
}

// DeleteUsuario : deleta a Usuario pelo seu id
func (u usuarioController) DeleteUsuario(c *gin.Context) {
	log.Print("[UsuarioController]...Deletando usuário")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	usuarioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao deletar usuário: %s", erro.Error())})
		return
	}

	usuarioTokenID, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao deletar usuário: %s", erro.Error())})
		return
	}
	if usuarioID != usuarioTokenID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Não é possível deletar um usuário que não o seu"})
		return
	}
	utensilio, erro := u.usuarioService.FindById(usuarioID, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar usuário: %s", erro.Error())})
		return
	}
	if utensilio.ID == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Não é possível deletar um usuário para uma empresa que não a sua"})
		return
	}

	erro = u.usuarioService.WithTrx(txHandle).Delete(usuarioID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar usuário: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindUsuarioById : busca a Usuario pelo seu id
func (u usuarioController) FindUsuarioById(c *gin.Context) {
	log.Print("[UsuarioController]...Buscando usuário por id")

	usuarioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar usuário: %s", erro.Error())})
		return
	}

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao buscar usuário: %s", erro.Error())})
		return
	}

	usuario, erro := u.usuarioService.FindById(usuarioID, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar usuário: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usuario})
}

// UsuarioController : busca todas as Usuarios de acordo com os parâmetros passados
func (u usuarioController) GetAllUsuarios(c *gin.Context) {
	log.Print("[UsuarioController]...Buscando todas as usuários")

	var empresaId uint64
	var erro error

	nome := c.Query("nome")
	if c.Query("empresaId") != "" {
		empresaId, erro = strconv.ParseUint(c.Query("empresaId"), 10, 64)
		if erro != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar usuários: %s", erro.Error())})
			return
		}
	}

	usuario := model.Usuario{
		Nome:      nome,
		EmpresaID: empresaId,
	}

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao buscar usuários: %s", erro.Error())})
		return
	}

	usuarios, erro := u.usuarioService.GetAll(usuario, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar usuários: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usuarios})
}
