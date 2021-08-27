package repository

import (
	"cookbook/src/constants"
	"cookbook/src/model"
	"errors"
	"fmt"
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
	FindById(uint64, uint64) (model.Receita, error)
	GetAll(string, constants.Categoria, uint64, uint64) ([]model.Receita, error)
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

	erro := r.DB.Omit("Utensilios.*").Create(&receita).Error

	return receita, erro
}

// Update -> atualiza uma receita no banco de dados
func (r receitaRepository) Update(receita model.Receita) (model.Receita, error) {
	log.Print("[ReceitaRepository]...Update")

	erro := r.DB.Where("receita_id = ?", receita.ID).Delete(&model.IngredienteReceita{}).Error
	if erro != nil {
		return receita, erro
	}

	utensilios := receita.Utensilios
	erro = r.DB.Model(&receita).Association("Utensilios").Clear()
	if erro != nil {
		return receita, erro
	}
	receita.Utensilios = utensilios

	erro = r.DB.Omit("Utensilios.*").Updates(&receita).Error

	return receita, erro
}

// Delete : deleta uma receita no banco de dados
func (r receitaRepository) Delete(id uint64) error {
	log.Print("[ReceitaRepository]...Delete")

	erro := r.DB.Delete(&model.Receita{}, id).Error

	return erro
}

// FindById -> busca uma receita pelo id no banco de dados
func (r receitaRepository) FindById(receitaID uint64, empresaID uint64) (receita model.Receita, erro error) {
	log.Print("[ReceitaRepository]...FindById")

	erro = r.DB.Preload("Ingredientes").Preload("Utensilios").Joins("JOIN usuarios ON receita.usuario_id = usuarios.id AND usuarios.empresa_id = ?", empresaID).Where("receita.id = ?", receitaID).First(&receita).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return receita, nil
	}

	return receita, erro
}

// GetAll -> busca todos as receitas no banco de dados de acordo com os parâmetros passados
func (r receitaRepository) GetAll(descricao string, categoria constants.Categoria, empresaID uint64, usuarioID uint64) (receitas []model.Receita, erro error) {
	log.Print("[ReceitaRepository]...GetAll")

	descricaoBusca := fmt.Sprintf("%%%s%%", descricao)

	if categoria != 0 {
		if usuarioID != 0 {
			erro = r.DB.Where("descricao LIKE ? AND categoria = ? AND usuario_id = ?", descricaoBusca, categoria, usuarioID).Find(&receitas).Error
		} else if empresaID != 0 {
			erro = r.DB.Joins("JOIN usuarios ON receita.usuario_id = usuarios.id AND usuarios.empresa_id = ?", empresaID).Where("descricao LIKE ? AND categoria = ?", descricaoBusca, categoria).Find(&receitas).Error
		} else {
			erro = r.DB.Where("descricao LIKE ? AND categoria = ?", descricaoBusca, categoria).Find(&receitas).Error
		}
	} else {
		if usuarioID != 0 {
			erro = r.DB.Where("descricao LIKE ? AND usuario_id = ?", descricaoBusca, usuarioID).Find(&receitas).Error
		} else if empresaID != 0 {
			erro = r.DB.Joins("JOIN usuarios ON receita.usuario_id = usuarios.id AND usuarios.empresa_id = ?", empresaID).Where("descricao LIKE ?", descricaoBusca).Find(&receitas).Error
		} else {
			erro = r.DB.Where("descricao LIKE ?", descricaoBusca).Find(&receitas).Error
		}
	}
	return receitas, erro
}
