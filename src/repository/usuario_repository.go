package repository

import (
	"cookbook/src/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type usuarioRepository struct {
	DB *gorm.DB
}

// UsuarioRepository : Representa o contrato de usuarios repository
type UsuarioRepository interface {
	WithTrx(*gorm.DB) UsuarioRepository
	Save(model.Usuario) (model.Usuario, error)
	Update(model.Usuario) (model.Usuario, error)
	Delete(uint64) error
	FindById(uint64, uint64) (model.Usuario, error)
	FindByUsername(string) (model.Usuario, error)
	GetAll(model.Usuario, uint64) ([]model.Usuario, error)
}

// NewUsuarioRepository -> retorna um novo usuario repository
func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return usuarioRepository{
		DB: db,
	}
}

// WithTrx : inicia uma transação para a ação que sera utilizada
func (u usuarioRepository) WithTrx(trxHandle *gorm.DB) UsuarioRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle

	return u
}

// Save -> salva um novo usuario no banco de dados
func (u usuarioRepository) Save(usuario model.Usuario) (model.Usuario, error) {
	log.Print("[usuarioRepository]...Save")

	erro := u.DB.Create(&usuario).Error

	return usuario, erro
}

// Update -> atualiza um usuario no banco de dados
func (u usuarioRepository) Update(usuario model.Usuario) (model.Usuario, error) {
	log.Print("[usuarioRepository]...Update")

	erro := u.DB.Save(&usuario).Error

	return usuario, erro
}

// Delete : deleta um usuario no banco de dados
func (u usuarioRepository) Delete(id uint64) error {
	log.Print("[usuarioRepository]...Delete")

	erro := u.DB.Delete(&model.Usuario{}, id).Error

	return erro
}

// FindById -> busca um usuario pelo id no banco de dados
func (u usuarioRepository) FindById(usuarioID uint64, empresaID uint64) (usuario model.Usuario, erro error) {
	log.Print("[usuarioRepository]...FindById")

	erro = u.DB.Where("id = ?", usuarioID).First(&usuario).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return usuario, nil
	}

	return usuario, erro
}

// FindByUsername -> busca o usuario no banco de dados pelo nick
func (u usuarioRepository) FindByUsername(username string) (usuario model.Usuario, erro error) {
	log.Print("[usuarioRepository]...FindByNick")

	erro = u.DB.Where("username = ?", username).Find(&usuario).Error

	return usuario, erro
}

// GetAll -> busca todos os usuarios no banco de dados que correspondem a descrição passada
func (u usuarioRepository) GetAll(usuario model.Usuario, empresaID uint64) (usuarios []model.Usuario, erro error) {
	log.Print("[usuarioRepository]...GetAll")

	usernameBusca := fmt.Sprintf("%%%s%%", usuario.Username)

	if usuario.EmpresaID != 0 {
		erro = u.DB.Where("username LIKE ? and empresa_id = ?", usernameBusca, empresaID).Find(&usuarios).Error
	} else {
		erro = u.DB.Where("username LIKE ? and empresa_id = ?", usernameBusca, empresaID).Find(&usuarios).Error
	}

	return usuarios, erro
}
