package repository

import (
	"cookbook/src/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type ingredienteRepository struct {
	DB *gorm.DB
}

// IngredienteRepository : Representa o contrato de ingredientes repository
type IngredienteRepository interface {
	Migrate() error
	WithTrx(*gorm.DB) IngredienteRepository
	Save(model.Ingrediente) (model.Ingrediente, error)
	Update(model.Ingrediente) (model.Ingrediente, error)
	Delete(uint64) error
	FindById(uint64) (model.Ingrediente, error)
	GetAll(string) ([]model.Ingrediente, error)
}

// NewingredienteRepository -> retorna um novo ingrediente repository
func NewIngredienteRepository(db *gorm.DB) IngredienteRepository {
	return ingredienteRepository{
		DB: db,
	}
}

// Migrate : Irá criar a tabela de Ingrediente no banco de dados
func (i ingredienteRepository) Migrate() error {
	log.Print("[IngredienteRepository]...Migrate")
	return i.DB.AutoMigrate(&model.Ingrediente{})
}

// WithTrx : Inicia uma transação para a ação que sera utilizada
func (i ingredienteRepository) WithTrx(trxHandle *gorm.DB) IngredienteRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return i
	}
	i.DB = trxHandle

	return i
}

// Save : Salva um novo ingrediente no banco de dados
func (i ingredienteRepository) Save(ingrediente model.Ingrediente) (model.Ingrediente, error) {
	log.Print("[IngredienteRepository]...Save")

	erro := i.DB.Create(&ingrediente).Error

	return ingrediente, erro
}

// Update : Atualiza um ingrediente no banco de dados
func (i ingredienteRepository) Update(ingrediente model.Ingrediente) (model.Ingrediente, error) {
	log.Print("[IngredienteRepository]...Update")

	erro := i.DB.Save(&ingrediente).Error

	return ingrediente, erro
}

// Delete : Deleta um ingrediente no banco de dados
func (i ingredienteRepository) Delete(id uint64) error {
	log.Print("[IngredienteRepository]...Delete")

	erro := i.DB.Delete(&model.Ingrediente{}, id).Error

	return erro
}

// FindById : Busca um ingrediente pelo id no banco de dados
func (i ingredienteRepository) FindById(id uint64) (ingrediente model.Ingrediente, erro error) {
	log.Print("[IngredienteRepository]...FindById")

	erro = i.DB.Where("id = ?", id).First(&ingrediente).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return ingrediente, nil
	}

	return ingrediente, erro
}

// GetAll : Busca os ingrediente no banco de dados que correspondem a descrição passada
func (i ingredienteRepository) GetAll(descricao string) (ingredientes []model.Ingrediente, erro error) {
	log.Print("[IngredienteRepository]...GetAll")

	descricaoBusca := fmt.Sprintf("%%%s%%", descricao)
	erro = i.DB.Where("descricao LIKE ?", descricaoBusca).Find(&ingredientes).Error

	return ingredientes, erro
}
