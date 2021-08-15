package repository

import (
	"cookbook/src/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type utensilioRepository struct {
	DB *gorm.DB
}

// UtensilioRepository : Representa o contrato de Utensilios repository
type UtensilioRepository interface {
	WithTrx(*gorm.DB) UtensilioRepository
	Save(model.Utensilio) (model.Utensilio, error)
	Update(model.Utensilio) (model.Utensilio, error)
	Delete(uint64) error
	FindById(uint64, uint64) (model.Utensilio, error)
	GetAll(string, uint64) ([]model.Utensilio, error)
}

// NewUtensilioRepository -> retorna um novo Utensilio repository
func NewUtensilioRepository(db *gorm.DB) UtensilioRepository {
	return utensilioRepository{
		DB: db,
	}
}

// WithTrx : inicia uma transação para a ação que sera utilizada
func (u utensilioRepository) WithTrx(trxHandle *gorm.DB) UtensilioRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle

	return u
}

// Save -> salva um novo Utensilio no banco de dados
func (u utensilioRepository) Save(Utensilio model.Utensilio) (model.Utensilio, error) {
	log.Print("[UtensilioRepository]...Save")

	erro := u.DB.Create(&Utensilio).Error

	return Utensilio, erro
}

// Update -> atualiza um Utensilio no banco de dados
func (u utensilioRepository) Update(Utensilio model.Utensilio) (model.Utensilio, error) {
	log.Print("[UtensilioRepository]...Update")

	erro := u.DB.Save(&Utensilio).Error

	return Utensilio, erro
}

// Delete : deleta um Utensilio no banco de dados
func (u utensilioRepository) Delete(id uint64) error {
	log.Print("[UtensilioRepository]...Delete")

	erro := u.DB.Delete(&model.Utensilio{}, id).Error

	return erro
}

// FindById -> busca um Utensilio pelo id no banco de dados
func (u utensilioRepository) FindById(utensilioID uint64, empresaID uint64) (Utensilio model.Utensilio, erro error) {
	log.Print("[UtensilioRepository]...FindById")

	erro = u.DB.Where("id = ? and empresa_id", utensilioID, empresaID).First(&Utensilio).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return Utensilio, nil
	}

	return Utensilio, erro
}

// GetAll -> busca os Utensilio no banco de dados que correspondem a descrição passada
func (u utensilioRepository) GetAll(descricao string, empresaID uint64) (Utensilios []model.Utensilio, erro error) {
	log.Print("[UtensilioRepository]...GetAll")

	descricaoBusca := fmt.Sprintf("%%%s%%", descricao)
	erro = u.DB.Where("descricao LIKE ? and empresa_id = ?", descricaoBusca, empresaID).Find(&Utensilios).Error

	return Utensilios, erro
}
