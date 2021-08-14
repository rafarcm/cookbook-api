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

type empresaController struct {
	empresaService service.EmpresaService
}

// EmpresaController : representa o contrato de EmpresaController
type EmpresaController interface {
	AddEmpresa(*gin.Context)
	UpdateEmpresa(*gin.Context)
	DeleteEmpresa(*gin.Context)
	FindEmpresaById(*gin.Context)
	GetAllEmpresas(c *gin.Context)
}

//NewEmpresaController -> retorna um EmpresaController
func NewEmpresaController(r service.EmpresaService) EmpresaController {
	return empresaController{
		empresaService: r,
	}
}

// AddEmpresa: adiciona uma nova Empresa
func (e empresaController) AddEmpresa(c *gin.Context) {
	log.Print("[EmpresaController]...Adicionando Empresa")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	var Empresa model.Empresa
	if erro := c.ShouldBindJSON(&Empresa); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	Empresa, erro := e.empresaService.WithTrx(txHandle).Save(Empresa)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar Empresa: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": Empresa})
}

// UpdateEmpresa : atualiza a Empresa pelo seu id
func (e empresaController) UpdateEmpresa(c *gin.Context) {
	log.Print("[EmpresaController]...Atualizando Empresa")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	EmpresaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar Empresa: %s", erro.Error())})
		return
	}

	var Empresa model.Empresa
	if erro := c.ShouldBindJSON(&Empresa); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar Empresa: %s", erro.Error())})
		return
	}

	Empresa.ID = EmpresaID
	Empresa, erro = e.empresaService.WithTrx(txHandle).Update(Empresa)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao atualizar Empresa: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Empresa})
}

// DeleteEmpresa : deleta a Empresa pelo seu id
func (e empresaController) DeleteEmpresa(c *gin.Context) {
	log.Print("[EmpresaController]...Deletando Empresa")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	EmpresaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao deletar Empresa: %s", erro.Error())})
		return
	}

	erro = e.empresaService.WithTrx(txHandle).Delete(EmpresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar Empresa: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindEmpresaById : busca a Empresa pelo seu id
func (e empresaController) FindEmpresaById(c *gin.Context) {
	log.Print("[EmpresaController]...Buscando Empresa por id")

	EmpresaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar Empresa: %s", erro.Error())})
		return
	}

	Empresa, erro := e.empresaService.FindById(EmpresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar Empresa: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Empresa})
}

// EmpresaController : busca todas as Empresas de acordo com os parâmetros passados
func (e empresaController) GetAllEmpresas(c *gin.Context) {
	log.Print("[EmpresaController]...Buscando todas as Empresas")

	descricao := c.Query("descricao")

	empresas, erro := e.empresaService.GetAll(descricao)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar Empresas: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": empresas})
}
