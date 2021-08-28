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
	UpdateSenha(uint64, string) error
	Delete(uint64) error
	FindById(uint64, uint64) (model.Usuario, error)
	FindByEmail(string) (model.Usuario, error)
	GetAll(string, uint64) ([]model.Usuario, error)
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

// UpdateSenha : atualiza a senha de um usuario no banco de dados
func (u usuarioRepository) UpdateSenha(id uint64, senha string) error {
	log.Print("[usuarioRepository]...UpdateSenha")

	erro := u.DB.Model(model.Usuario{}).Where("id = ?", id).Update("senha", senha).Error

	return erro
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

	erro = u.DB.Where("id = ? AND empresa_id = ?", usuarioID, empresaID).First(&usuario).Error

	if erro != nil && errors.Is(erro, gorm.ErrRecordNotFound) {
		return usuario, nil
	}

	return usuario, erro
}

// FindByEmail -> busca o usuario no banco de dados pelo email
func (u usuarioRepository) FindByEmail(email string) (usuario model.Usuario, erro error) {
	log.Print("[usuarioRepository]...FindByEmail")

	erro = u.DB.Where("email = ?", email).Find(&usuario).Error

	return usuario, erro
}

// GetAll -> busca todos os usuarios no banco de dados que correspondem ao email passado
func (u usuarioRepository) GetAll(nome string, empresaID uint64) (usuarios []model.Usuario, erro error) {
	log.Print("[usuarioRepository]...GetAll")

	nomeBusca := fmt.Sprintf("%%%s%%", nome)

	if empresaID != 0 {
		erro = u.DB.Where("nome LIKE ? AND empresa_id = ?", nomeBusca, empresaID).Find(&usuarios).Error
	} else {
		erro = u.DB.Where("nome LIKE ?", nomeBusca).Find(&usuarios).Error
	}

	return usuarios, erro
}
