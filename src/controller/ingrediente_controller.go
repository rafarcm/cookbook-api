package controller

import (
	"cookbook/src/model"
	"cookbook/src/service"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ingredienteController struct {
	ingredienteService service.IngredienteService
}

// IngredienteController : Representa o contrato de ingredientes controller
type IngredienteController interface {
	AddIngrediente(*gin.Context)
	UpdateIngrediente(*gin.Context)
	DeleteIngrediente(*gin.Context)
	FindIngredienteById(*gin.Context)
	GetAllIngredientes(*gin.Context)
}

//NewIngredienteController -> retorna um novo ingrediente controller
func NewIngredienteController(s service.IngredienteService) IngredienteController {
	return ingredienteController{
		ingredienteService: s,
	}
}

// AddIngrediente : Adiciona um novo ingrediente
func (i ingredienteController) AddIngrediente(c *gin.Context) {
	log.Print("[IngredienteController]...Adicionando ingrediente")

	var ingrediente model.Ingrediente
	if erro := c.ShouldBindJSON(&ingrediente); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	ingrediente, erro := i.ingredienteService.Save(ingrediente)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar ingrediente"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": ingrediente})
}

// FindIngredienteById : Busca o ingrediente pelo seu id
func (i ingredienteController) UpdateIngrediente(c *gin.Context) {
	log.Print("[IngredienteController]...Atualizando ingrediente")

	ingredienteID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	var ingrediente model.Ingrediente
	if erro := c.ShouldBindJSON(&ingrediente); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	ingrediente, erro = i.ingredienteService.Update(ingredienteID, ingrediente)
	if erro != nil && !errors.Is(erro, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o ingrediente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingrediente})
}

// DeleteIngrediente : Deleta o ingrediente pelo seu id
func (i ingredienteController) DeleteIngrediente(c *gin.Context) {
	log.Print("[IngredienteController]...Deletando ingrediente")

	ingredienteID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	erro = i.ingredienteService.Delete(ingredienteID)
	if erro != nil && !errors.Is(erro, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o ingrediente"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindIngredienteById : Busca o ingrediente pelo seu id
func (i ingredienteController) FindIngredienteById(c *gin.Context) {
	log.Print("[IngredienteController]...Buscando ingrediente por id")

	ingredienteID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	ingrediente, erro := i.ingredienteService.FindById(ingredienteID)
	if erro != nil && !errors.Is(erro, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o ingrediente pelo id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingrediente})
}

// GetAllIngredientes : Busca todos os ingredientes pela descrição desejada
func (i ingredienteController) GetAllIngredientes(c *gin.Context) {
	log.Print("[IngredienteController]...Buscando todos os ingredientes")

	descricao := c.Query("descricao")

	ingredientes, erro := i.ingredienteService.GetAll(descricao)
	if erro != nil && !errors.Is(erro, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar os ingredientes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingredientes})
}
