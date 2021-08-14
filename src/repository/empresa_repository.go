package repository

import (
	"cookbook/src/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type empresaRepository struct {
	DB *gorm.DB
}

// EmpresaRepository : Representa o contrato de Empresas repository
type EmpresaRepository interface {
	WithTrx(*gorm.DB) EmpresaRepository
	Save(model.Empresa) (model.Empresa, error)
	Update(model.Empresa) (model.Empresa, error)
	Delete(uint64) error
	FindById(uint64) (model.Empresa, error)
	GetAll(string) ([]model.Empresa, error)
}

// NewEmpresaRepository -> retorna um novo Empresa repository
func NewEmpresaRepository(db *gorm.DB) EmpresaRepository {
	return empresaRepository{
		DB: db,
	}
}

// WithTrx : inicia uma transação para a ação que sera utilizada
func (e empresaRepository) WithTrx(trxHandle *gorm.DB) EmpresaRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return e
	}
	e.DB = trxHandle

	return e
}

// Save -> salva um novo Empresa no banco de dados
func (e empresaRepository) Save(empresa model.Empresa) (model.Empresa, error) {
	log.Print("[EmpresaRepository]...Save")

	erro := e.DB.Create(&empresa).Error

	return empresa, erro
}

// Update -> atualiza um Empresa no banco de dados
func (e empresaRepository) Update(empresa model.Empresa) (model.Empresa, error) {
	log.Print("[EmpresaRepository]...Update")

	erro := e.DB.Save(&empresa).Error

	return empresa, erro
}

// Delete : deleta um Empresa no banco de dados
func (e empresaRepository) Delete(id uint64) error {
	log.Print("[EmpresaRepository]...Delete")

	erro := e.DB.Delete(&model.Empresa{}, id).Error

	return erro
}

// FindById -> busca um Empresa pelo id no banco de dados
func (e empresaRepository) FindById(id uint64) (empresa model.Empresa, erro error) {
	log.Print("[EmpresaRepository]...FindById")

	erro = e.DB.Where("id = ?", id).First(&empresa).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return empresa, nil
	}

	return empresa, erro
}

// GetAll -> busca todos os Empresas no banco de dados que correspondem a descrição passada
func (e empresaRepository) GetAll(nome string) (empresas []model.Empresa, erro error) {
	log.Print("[EmpresaRepository]...GetAll")

	descricaoBusca := fmt.Sprintf("%%%s%%", nome)
	erro = e.DB.Where("nome LIKE ?", descricaoBusca).Find(&empresas).Error

	return empresas, erro
}
