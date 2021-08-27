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
	GetAll(string, uint64, uint64) ([]model.Utensilio, error)
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

	erro = u.DB.Joins("JOIN usuarios ON utensilios.usuario_id = usuarios.id AND usuarios.empresa_id = ?", empresaID).Where("utensilios.id = ?", utensilioID).First(&Utensilio).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return Utensilio, nil
	}

	return Utensilio, erro
}

// GetAll -> busca os Utensilio no banco de dados que correspondem a descrição passada
func (u utensilioRepository) GetAll(descricao string, empresaID uint64, usuarioID uint64) (utensilios []model.Utensilio, erro error) {
	log.Print("[UtensilioRepository]...GetAll")

	descricaoBusca := fmt.Sprintf("%%%s%%", descricao)
	if usuarioID != 0 {
		erro = u.DB.Where("descricao LIKE ? AND usuario_id = ?", descricaoBusca, usuarioID).Find(&utensilios).Error
	} else if empresaID != 0 {
		erro = u.DB.Joins("JOIN usuarios ON utensilios.usuario_id = usuarios.id AND usuarios.empresa_id = ?", empresaID).Where("descricao LIKE ?", descricaoBusca).Find(&utensilios).Error
	} else {
		erro = u.DB.Where("descricao LIKE", descricaoBusca).Find(&utensilios).Error
	}

	return utensilios, erro
}
