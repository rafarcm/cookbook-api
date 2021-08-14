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

	var Utensilio model.Utensilio
	if erro := c.ShouldBindJSON(&Utensilio); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	utensilio, erro := u.utensilioService.WithTrx(txHandle).Save(Utensilio)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar utensílio: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": utensilio})
}

// UpdateUtensilio : atualiza a Utensilio pelo seu id
func (u utensilioController) UpdateUtensilio(c *gin.Context) {
	log.Print("[UtensilioController]...Atualizando utensílio")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	UtensilioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	var Utensilio model.Utensilio
	if erro := c.ShouldBindJSON(&Utensilio); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	Utensilio.ID = UtensilioID
	Utensilio, erro = u.utensilioService.WithTrx(txHandle).Update(Utensilio)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao atualizar utensílio: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Utensilio})
}

// DeleteUtensilio : deleta a Utensilio pelo seu id
func (u utensilioController) DeleteUtensilio(c *gin.Context) {
	log.Print("[UtensilioController]...Deletando utensílio")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	UtensilioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao deletar utensílio: %s", erro.Error())})
		return
	}

	erro = u.utensilioService.WithTrx(txHandle).Delete(UtensilioID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar utensílio: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindUtensilioById : busca a Utensilio pelo seu id
func (u utensilioController) FindUtensilioById(c *gin.Context) {
	log.Print("[UtensilioController]...Buscando utensílio por id")

	UtensilioID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar utensílio: %s", erro.Error())})
		return
	}

	Utensilio, erro := u.utensilioService.FindById(UtensilioID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar utensílio: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Utensilio})
}

// GetAllUtensilios : busca todos os utensilios pela descrição desejada
func (u utensilioController) GetAllUtensilios(c *gin.Context) {
	log.Print("[utensilioController]...Buscando todos os utensilios")

	descricao := c.Query("descricao")

	utensilios, erro := u.utensilioService.GetAll(descricao)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar utensilios: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": utensilios})
}
