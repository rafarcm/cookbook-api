package service

import (
	"cookbook/src/model"
	"cookbook/src/repository"
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
	Update(model.Usuario) (model.Usuario, error)
	Delete(uint64) error
	FindById(uint64) (model.Usuario, error)
	GetAll(model.Usuario) ([]model.Usuario, error)
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
func (u usuarioService) Save(Usuario model.Usuario) (model.Usuario, error) {
	Usuario.CriadoEm = time.Now()
	Usuario.AtualizadoEm = Usuario.CriadoEm

	return u.usuarioRepository.Save(Usuario)
}

// Update -> atualiza a Usuario e a retorna
func (u usuarioService) Update(Usuario model.Usuario) (model.Usuario, error) {
	UsuarioBanco, erro := u.usuarioRepository.FindById(Usuario.ID)
	if erro != nil {
		return model.Usuario{}, erro
	}

	Usuario.CriadoEm = UsuarioBanco.CriadoEm
	Usuario.AtualizadoEm = time.Now()

	return u.usuarioRepository.Update(Usuario)
}

// Delete -> exclui uma Usuario com o id passado
func (u usuarioService) Delete(id uint64) error {
	return u.usuarioRepository.Delete(id)
}

// FindById -> retorna a Usuario com o id passado
func (u usuarioService) FindById(id uint64) (model.Usuario, error) {
	return u.usuarioRepository.FindById(id)
}

// GetAll -> retorna todos os Usuarioss cadastrados que contém a descrição desejada
func (u usuarioService) GetAll(usuario model.Usuario) ([]model.Usuario, error) {
	return u.usuarioRepository.GetAll(usuario)
}
