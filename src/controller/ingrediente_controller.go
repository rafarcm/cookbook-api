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

type ingredienteController struct {
	ingredienteService service.IngredienteService
}

// IngredienteController : representa o contrato de IngredienteController
type IngredienteController interface {
	AddIngrediente(*gin.Context)
	UpdateIngrediente(*gin.Context)
	DeleteIngrediente(*gin.Context)
	FindIngredienteById(*gin.Context)
	GetAllIngredientes(*gin.Context)
}

//NewIngredienteController -> retorna um novo IngredienteController
func NewIngredienteController(s service.IngredienteService) IngredienteController {
	return ingredienteController{
		ingredienteService: s,
	}
}

// AddIngrediente : adiciona um novo ingrediente
func (i ingredienteController) AddIngrediente(c *gin.Context) {
	log.Print("[IngredienteController]...Adicionando ingrediente")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	var ingrediente model.Ingrediente
	if erro := c.ShouldBindJSON(&ingrediente); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao adicionar ingrediente: %s", erro.Error())})
		return
	}

	ingrediente, erro := i.ingredienteService.WithTrx(txHandle).Save(ingrediente)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar ingrediente: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": ingrediente})
}

// UpdateIngrediente : atualiza o ingrediente pelo seu id
func (i ingredienteController) UpdateIngrediente(c *gin.Context) {
	log.Print("[IngredienteController]...Atualizando ingrediente")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	ingredienteID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar ingrediente: %s", erro.Error())})
		return
	}

	var ingrediente model.Ingrediente
	if erro := c.ShouldBindJSON(&ingrediente); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar ingrediente: %s", erro.Error())})
		return
	}

	ingrediente.ID = ingredienteID
	ingrediente, erro = i.ingredienteService.WithTrx(txHandle).Update(ingrediente)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao atualizar ingrediente: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingrediente})
}

// DeleteIngrediente : deleta o ingrediente pelo seu id
func (i ingredienteController) DeleteIngrediente(c *gin.Context) {
	log.Print("[IngredienteController]...Deletando ingrediente")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	ingredienteID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao deletar ingrediente: %s", erro.Error())})
		return
	}

	erro = i.ingredienteService.WithTrx(txHandle).Delete(ingredienteID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar ingrediente: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindIngredienteById : busca o ingrediente pelo seu id
func (i ingredienteController) FindIngredienteById(c *gin.Context) {
	log.Print("[IngredienteController]...Buscando ingrediente por id")

	ingredienteID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar ingrediente: %s", erro.Error())})
		return
	}

	ingrediente, erro := i.ingredienteService.FindById(ingredienteID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar ingrediente: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingrediente})
}

// GetAllIngredientes : busca todos os ingredientes pela descrição desejada
func (i ingredienteController) GetAllIngredientes(c *gin.Context) {
	log.Print("[IngredienteController]...Buscando todos os ingredientes")

	descricao := c.Query("descricao")

	ingredientes, erro := i.ingredienteService.GetAll(descricao)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar ingredientes: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingredientes})
}
