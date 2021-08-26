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

	var empresa model.Empresa
	if erro := c.ShouldBindJSON(&empresa); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	for i := range empresa.Usuarios {
		senhaComHash, erro := security.Hash(empresa.Usuarios[i].Senha)
		if erro != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar empresa: %s", erro.Error())})
			return
		}
		empresa.Usuarios[i].Senha = string(senhaComHash)
	}

	empresa, erro := e.empresaService.WithTrx(txHandle).Save(empresa)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao adicionar empresa: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": empresa})
}

// UpdateEmpresa : atualiza a Empresa pelo seu id
func (e empresaController) UpdateEmpresa(c *gin.Context) {
	log.Print("[EmpresaController]...Atualizando Empresa")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	empresaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar Empresa: %s", erro.Error())})
		return
	}

	_, empresaTokenID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao atualizar empresa: %s", erro.Error())})
		return
	}
	if empresaTokenID != empresaID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Não é possível atualizar uma empresa que não a sua"})
		return
	}

	var empresa model.Empresa
	if erro := c.ShouldBindJSON(&empresa); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao atualizar Empresa: %s", erro.Error())})
		return
	}

	empresa.ID = empresaID
	empresa, erro = e.empresaService.WithTrx(txHandle).Update(empresa)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao atualizar Empresa: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": empresa})
}

// DeleteEmpresa : deleta a Empresa pelo seu id
func (e empresaController) DeleteEmpresa(c *gin.Context) {
	log.Print("[EmpresaController]...Deletando Empresa")

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	empresaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao deletar Empresa: %s", erro.Error())})
		return
	}

	_, empresaTokenID, erro := authentication.ExtrairIDs(c)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Erro ao deletar empresa: %s", erro.Error())})
		return
	}
	if empresaTokenID != empresaID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Não é possível deletar uma empresa que não a sua"})
		return
	}

	erro = e.empresaService.WithTrx(txHandle).Delete(empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao deletar Empresa: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindEmpresaById : busca a Empresa pelo seu id
func (e empresaController) FindEmpresaById(c *gin.Context) {
	log.Print("[EmpresaController]...Buscando Empresa por id")

	empresaID, erro := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao buscar Empresa: %s", erro.Error())})
		return
	}

	empresa, erro := e.empresaService.FindById(empresaID)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar Empresa: %s", erro.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": empresa})
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
