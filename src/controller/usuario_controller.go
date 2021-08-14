package controller

import (
	"cookbook/src/model"
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
	log.Print("[UsuarioController]...Adicionando Usuario")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	var Usuario model.Usuario
	if erro := c.ShouldBindJSON(&Usuario); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	Usuario, erro := u.usuarioService.WithTrx(txHandle).Save(Usuario)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar Usuario: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": Usuario})
}

// UpdateUsuario : atualiza a Usuario pelo seu id
func (u usuarioController) UpdateUsuario(c *gin.Context) {
	log.Print("[UsuarioController]...Atualizando Usuario")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	UsuarioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar Usuario: %s", erro.Error())})
		return
	}

	var Usuario model.Usuario
	if erro := c.ShouldBindJSON(&Usuario); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar Usuario: %s", erro.Error())})
		return
	}

	Usuario.ID = UsuarioID
	Usuario, erro = u.usuarioService.WithTrx(txHandle).Update(Usuario)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao atualizar Usuario: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Usuario})
}

// DeleteUsuario : deleta a Usuario pelo seu id
func (u usuarioController) DeleteUsuario(c *gin.Context) {
	log.Print("[UsuarioController]...Deletando Usuario")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	UsuarioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao deletar Usuario: %s", erro.Error())})
		return
	}

	erro = u.usuarioService.WithTrx(txHandle).Delete(UsuarioID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar Usuario: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindUsuarioById : busca a Usuario pelo seu id
func (u usuarioController) FindUsuarioById(c *gin.Context) {
	log.Print("[UsuarioController]...Buscando Usuario por id")

	UsuarioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar Usuario: %s", erro.Error())})
		return
	}

	Usuario, erro := u.usuarioService.FindById(UsuarioID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar Usuario: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Usuario})
}

// UsuarioController : busca todas as Usuarios de acordo com os parâmetros passados
func (u usuarioController) GetAllUsuarios(c *gin.Context) {
	log.Print("[UsuarioController]...Buscando todas as Usuarios")

	var empresaId uint64
	var erro error

	nick := c.Query("nick")
	if c.Query("empresaId") != "" {
		empresaId, erro = strconv.ParseUint(c.Query("empresaId"), 10, 64)
		if erro != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar Usuarios: %s", erro.Error())})
			return
		}
	}

	usuario := model.Usuario{
		Nick:      nick,
		EmpresaID: empresaId,
	}

	usuarios, erro := u.usuarioService.GetAll(usuario)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar Usuarios: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usuarios})
}
