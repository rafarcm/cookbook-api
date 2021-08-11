package repository

import (
	"cookbook/src/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type receitaRepository struct {
	DB *gorm.DB
}

// ReceitaRepository : representa o contrato de receita repository
type ReceitaRepository interface {
	WithTrx(*gorm.DB) ReceitaRepository
	Save(model.Receita) (model.Receita, error)
	Update(model.Receita) (model.Receita, error)
	Delete(uint64) error
	FindById(uint64) (model.Receita, error)
}

// NewReceitaRepository -> retorna um novo receita repository
func NewReceitaRepository(db *gorm.DB) ReceitaRepository {
	return receitaRepository{
		DB: db,
	}
}

// WithTrx -> inicia uma transação para a ação que sera utilizada
func (r receitaRepository) WithTrx(trxHandle *gorm.DB) ReceitaRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return r
	}
	r.DB = trxHandle

	return r
}

// Save -> salva uma nova receita no banco de dados
func (r receitaRepository) Save(receita model.Receita) (model.Receita, error) {
	log.Print("[ReceitaRepository]...Save")

	erro := r.DB.Create(&receita).Error

	return receita, erro
}

// Update -> atualiza uma receita no banco de dados
func (r receitaRepository) Update(receita model.Receita) (model.Receita, error) {
	log.Print("[ReceitaRepository]...Update")

	erro := r.DB.Where("receita_id = ?", receita.ID).Delete(&model.IngredienteReceita{}).Error
	if erro != nil {
		return receita, erro
	}

	erro = r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&receita).Error

	return receita, erro
}

// Delete : deleta uma receita no banco de dados
func (r receitaRepository) Delete(id uint64) error {
	log.Print("[ReceitaRepository]...Delete")

	erro := r.DB.Delete(&model.Receita{}, id).Error

	return erro
}

// FindById -> busca uma receita pelo id no banco de dados
func (r receitaRepository) FindById(id uint64) (receita model.Receita, erro error) {
	log.Print("[ReceitaRepository]...FindById")

	erro = r.DB.Preload("Ingredientes").Where("id = ?", id).First(&receita).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return receita, nil
	}

	return receita, erro
}
