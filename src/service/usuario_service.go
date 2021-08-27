package service

import (
	"cookbook/src/model"
	"cookbook/src/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type usuarioService struct {
	usuarioRepository repository.UsuarioRepository
}

// UsuarioService : representa o contrato de Usuario service
type UsuarioService interface {
	WithTrx(*gorm.DB) UsuarioService
	Save(model.Usuario) (model.Usuario, error)
	Update(model.Usuario, uint64) (model.Usuario, error)
	Delete(uint64, uint64) error
	FindById(uint64, uint64) (model.Usuario, error)
	FindByEmail(string) (model.Usuario, error)
	GetAll(string, uint64) ([]model.Usuario, error)
}

// NewUsuarioService -> retorna um novo Usuario service
func NewUsuarioService(u repository.UsuarioRepository) UsuarioService {
	return usuarioService{
		usuarioRepository: u,
	}
}

// WithTrx : habilita repositório com transação
func (u usuarioService) WithTrx(trxHandle *gorm.DB) UsuarioService {
	u.usuarioRepository = u.usuarioRepository.WithTrx(trxHandle)
	return u
}

// Save -> salva uma nova Usuario e a retorna
func (u usuarioService) Save(usuario model.Usuario) (model.Usuario, error) {
	usuario.CriadoEm = time.Now()
	usuario.AtualizadoEm = usuario.CriadoEm

	return u.usuarioRepository.Save(usuario)
}

// Update -> atualiza a Usuario e a retorna
func (u usuarioService) Update(usuario model.Usuario, empresaID uint64) (model.Usuario, error) {
	usuarioBanco, erro := u.usuarioRepository.FindById(usuario.ID, empresaID)
	if erro != nil {
		return model.Usuario{}, erro
	}

	if usuarioBanco.ID == 0 {
		return model.Usuario{}, errors.New("não é possível alterar este usuário")
	}

	usuario.CriadoEm = usuarioBanco.CriadoEm
	usuario.AtualizadoEm = time.Now()
	usuario.Senha = usuarioBanco.Senha

	return u.usuarioRepository.Update(usuario)
}

// Delete -> exclui uma Usuario com o id passado
func (u usuarioService) Delete(id uint64, empresaID uint64) error {
	usuarioBanco, erro := u.usuarioRepository.FindById(id, empresaID)
	if erro != nil {
		return erro
	}

	if usuarioBanco.ID == 0 {
		return errors.New("não é possível deletar este usuário")
	}

	return u.usuarioRepository.Delete(id)
}

// FindById -> retorna a Usuario com o id passado
func (u usuarioService) FindById(receitaID uint64, empresaID uint64) (model.Usuario, error) {
	return u.usuarioRepository.FindById(receitaID, empresaID)
}

// FindByEmail -> retorna a Usuario com o email passado
func (u usuarioService) FindByEmail(username string) (model.Usuario, error) {
	return u.usuarioRepository.FindByEmail(username)
}

// GetAll -> retorna todos os Usuarioss cadastrados que contém a descrição desejada
func (u usuarioService) GetAll(nome string, empresaID uint64) ([]model.Usuario, error) {
	return u.usuarioRepository.GetAll(nome, empresaID)
}
