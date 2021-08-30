package controller

import (
	"cookbook/src/authentication"
	"cookbook/src/model"
	"cookbook/src/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type utensilioController struct {
	utensilioService service.UtensilioService
}

// UtensilioController : representa o contrato de UtensilioController
type UtensilioController interface {
	AddUtensilio(*gin.Context)
	UpdateUtensilio(*gin.Context)
	DeleteUtensilio(*gin.Context)
	FindUtensilioById(*gin.Context)
	GetAllUtensilios(*gin.Context)
}

//NewUtensilioController -> retorna um UtensilioController
func NewUtensilioController(u service.UtensilioService) UtensilioController {
	return utensilioController{
		utensilioService: u,
	}
}

// AddUtensilio: adiciona uma nova Utensilio
func (u utensilioController) AddUtensilio(c *gin.Context) {
	log.Print("[UtensilioController]...Adicionando utensílio")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	var utensilio model.Utensilio
	if erro := c.ShouldBindJSON(&utensilio); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	usuarioID, _, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao adicionar utensílio: %s", erro.Error())})
		return
	}
	utensilio.UsuarioID = usuarioID

	utensilio, erro = u.utensilioService.WithTrx(txHandle).Save(utensilio)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar utensílio: %s", erro.Error())})
		return
	}

	utensilio.AddLinks(c.Request.Host, model.UtensilioAddRel)

	c.JSON(http.StatusCreated, gin.H{"data": utensilio})
}

// UpdateUtensilio : atualiza a Utensilio pelo seu id
func (u utensilioController) UpdateUtensilio(c *gin.Context) {
	log.Print("[UtensilioController]...Atualizando utensílio")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	utensilioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	var utensilio model.Utensilio
	if erro := c.ShouldBindJSON(&utensilio); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	usuarioID, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	utensilio.UsuarioID = usuarioID
	utensilio.ID = utensilioID
	utensilio, erro = u.utensilioService.WithTrx(txHandle).Update(utensilio, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	utensilio.AddLinks(c.Request.Host, model.UtensilioUpdateRel)
	c.JSON(http.StatusOK, gin.H{"data": utensilio})
}

// DeleteUtensilio : deleta a Utensilio pelo seu id
func (u utensilioController) DeleteUtensilio(c *gin.Context) {
	log.Print("[UtensilioController]...Deletando utensílio")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	utensilioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao deletar utensílio: %s", erro.Error())})
		return
	}

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	utensilio, erro := u.utensilioService.FindById(utensilioID, empresaID)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	erro = u.utensilioService.WithTrx(txHandle).Delete(utensilioID, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar utensílio: %s", erro.Error())})
		return
	}

	utensilio.AddLinks(c.Request.Host, model.UtensilioDeleteRel)

	c.JSON(http.StatusOK, gin.H{"data": utensilio})
}

// FindUtensilioById : busca a Utensilio pelo seu id
func (u utensilioController) FindUtensilioById(c *gin.Context) {
	log.Print("[UtensilioController]...Buscando utensílio por id")

	utensilioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar utensílio: %s", erro.Error())})
		return
	}

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao buscar utensílios: %s", erro.Error())})
		return
	}

	utensilio, erro := u.utensilioService.FindById(utensilioID, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar utensílio: %s", erro.Error())})
		return
	}
	utensilio.AddLinks(c.Request.Host, model.UtensilioGetByIdRel)

	c.JSON(http.StatusOK, gin.H{"data": utensilio})
}

// GetAllUtensilios : busca todos os utensilios pela descrição desejada
func (u utensilioController) GetAllUtensilios(c *gin.Context) {
	log.Print("[utensilioController]...Buscando todos os utensilios")

	var usuarioID uint64 = 0
	var erro error

	descricao := c.Query("descricao")

	if c.Query("usuarioId") != "" {
		usuarioID, erro = strconv.ParseUint(c.Query("usuarioId"), 10, 64)
		if erro != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar utensílios: %s", erro.Error())})
			return
		}
	}

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao buscar utensílios: %s", erro.Error())})
		return
	}

	utensilios, erro := u.utensilioService.GetAll(descricao, empresaID, usuarioID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar utensilios: %s", erro.Error())})
		return
	}

	for i := range utensilios {
		utensilios[i].AddLinks(c.Request.Host, model.UtensilioGetAllRel)
	}

	c.JSON(http.StatusOK, gin.H{"data": utensilios})
}
