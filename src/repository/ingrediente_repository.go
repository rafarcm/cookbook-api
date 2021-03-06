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
	WithTrx(*gorm.DB) IngredienteRepository
	Save(model.Ingrediente) (model.Ingrediente, error)
	Update(model.Ingrediente) (model.Ingrediente, error)
	Delete(uint64) error
	FindById(uint64, uint64) (model.Ingrediente, error)
	GetAll(string, uint64, uint64) ([]model.Ingrediente, error)
}

// NewingredienteRepository -> retorna um novo ingrediente repository
func NewIngredienteRepository(db *gorm.DB) IngredienteRepository {
	return ingredienteRepository{
		DB: db,
	}
}

// WithTrx : inicia uma transação para a ação que sera utilizada
func (i ingredienteRepository) WithTrx(trxHandle *gorm.DB) IngredienteRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return i
	}
	i.DB = trxHandle

	return i
}

// Save -> salva um novo ingrediente no banco de dados
func (i ingredienteRepository) Save(ingrediente model.Ingrediente) (model.Ingrediente, error) {
	log.Print("[IngredienteRepository]...Save")

	erro := i.DB.Create(&ingrediente).Error

	return ingrediente, erro
}

// Update -> atualiza um ingrediente no banco de dados
func (i ingredienteRepository) Update(ingrediente model.Ingrediente) (model.Ingrediente, error) {
	log.Print("[IngredienteRepository]...Update")

	erro := i.DB.Save(&ingrediente).Error

	return ingrediente, erro
}

// Delete : deleta um ingrediente no banco de dados
func (i ingredienteRepository) Delete(id uint64) error {
	log.Print("[IngredienteRepository]...Delete")

	erro := i.DB.Delete(&model.Ingrediente{}, id).Error

	return erro
}

// FindById -> busca um ingrediente pelo id no banco de dados
func (i ingredienteRepository) FindById(ingredienteID uint64, empresaID uint64) (ingrediente model.Ingrediente, erro error) {
	log.Print("[IngredienteRepository]...FindById")

	erro = i.DB.Joins("JOIN usuarios ON ingredientes.usuario_id = usuarios.id AND usuarios.empresa_id = ?", empresaID).Where("ingredientes.id = ?", ingredienteID).First(&ingrediente).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return ingrediente, nil
	}

	return ingrediente, erro
}

// GetAll -> busca todos os ingredientes no banco de dados que correspondem a descrição passada
func (i ingredienteRepository) GetAll(descricao string, empresaID uint64, usuarioID uint64) (ingredientes []model.Ingrediente, erro error) {
	log.Print("[IngredienteRepository]...GetAll")

	descricaoBusca := fmt.Sprintf("%%%s%%", descricao)

	if usuarioID != 0 {
		erro = i.DB.Where("descricao LIKE ? AND usuario_id = ?", descricaoBusca, usuarioID).Find(&ingredientes).Error
	} else if empresaID != 0 {
		erro = i.DB.Joins("JOIN usuarios ON ingredientes.usuario_id = usuarios.id AND usuarios.empresa_id = ?", empresaID).Where("descricao LIKE ?", descricaoBusca).Find(&ingredientes).Error
	} else {
		erro = i.DB.Where("descricao LIKE", descricaoBusca).Find(&ingredientes).Error
	}

	return ingredientes, erro
}
