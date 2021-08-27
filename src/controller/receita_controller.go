package controller

import (
	"cookbook/src/authentication"
	"cookbook/src/constants"
	"cookbook/src/model"
	"cookbook/src/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type receitaController struct {
	receitaService service.ReceitaService
}

// ReceitaController : representa o contrato de ReceitaController
type ReceitaController interface {
	AddReceita(*gin.Context)
	UpdateReceita(*gin.Context)
	DeleteReceita(*gin.Context)
	FindReceitaById(*gin.Context)
	GetAllReceitas(c *gin.Context)
}

//NewReceitaController -> retorna um ReceitaController
func NewReceitaController(r service.ReceitaService) ReceitaController {
	return receitaController{
		receitaService: r,
	}
}

// AddReceita: adiciona uma nova receita
func (r receitaController) AddReceita(c *gin.Context) {
	log.Print("[ReceitaController]...Adicionando receita")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	var receita model.Receita
	if erro := c.ShouldBindJSON(&receita); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	usuarioID, _, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao adicionar receita: %s", erro.Error())})
		return
	}
	receita.UsuarioID = usuarioID

	receita, erro = r.receitaService.WithTrx(txHandle).Save(receita)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar receita: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": receita})
}

// UpdateReceita : atualiza a receita pelo seu id
func (r receitaController) UpdateReceita(c *gin.Context) {
	log.Print("[ReceitaController]...Atualizando receita")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	receitaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar receita: %s", erro.Error())})
		return
	}

	var receita model.Receita
	if erro := c.ShouldBindJSON(&receita); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar receita: %s", erro.Error())})
		return
	}

	usuarioID, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao atualizar receita: %s", erro.Error())})
		return
	}

	receita.UsuarioID = usuarioID
	receita.ID = receitaID
	receita, erro = r.receitaService.WithTrx(txHandle).Update(receita, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao atualizar receita: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": receita})
}

// DeleteReceita : deleta a receita pelo seu id
func (r receitaController) DeleteReceita(c *gin.Context) {
	log.Print("[ReceitaController]...Deletando receita")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	receitaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao deletar receita: %s", erro.Error())})
		return
	}

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao atualizar receita: %s", erro.Error())})
		return
	}

	erro = r.receitaService.WithTrx(txHandle).Delete(receitaID, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar receita: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindReceitaById : busca a receita pelo seu id
func (r receitaController) FindReceitaById(c *gin.Context) {
	log.Print("[ReceitaController]...Buscando receita por id")

	receitaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar receita: %s", erro.Error())})
		return
	}

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao buscar receita: %s", erro.Error())})
		return
	}

	receita, erro := r.receitaService.FindById(receitaID, empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar receita: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": receita})
}

// receitaController : busca todas as receitas de acordo com os parâmetros passados
func (r receitaController) GetAllReceitas(c *gin.Context) {
	log.Print("[ReceitaController]...Buscando todas as receitas")

	var usuarioID uint64 = 0
	var categoria uint64
	var erro error

	descricao := c.Query("descricao")

	if c.Query("usuarioId") != "" {
		usuarioID, erro = strconv.ParseUint(c.Query("usuarioId"), 10, 64)
		if erro != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar receitas: %s", erro.Error())})
			return
		}
	}

	if c.Query("categoria") != "" {
		categoria, erro = strconv.ParseUint(c.Query("categoria"), 10, 64)
		if erro != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar receitas: %s", erro.Error())})
			return
		}
	}

	_, empresaID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao buscar receitas: %s", erro.Error())})
		return
	}

	receitas, erro := r.receitaService.GetAll(descricao, constants.Categoria(categoria), empresaID, usuarioID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar receitas: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": receitas})
}
