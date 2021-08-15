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
	FindById(uint64, uint64) (model.Usuario, error)
	FindByUsername(string) (model.Usuario, error)
	GetAll(model.Usuario, uint64) ([]model.Usuario, error)
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
func (u usuarioService) Update(usuario model.Usuario) (model.Usuario, error) {
	usuarioBanco, erro := u.usuarioRepository.FindById(usuario.ID, usuario.EmpresaID)
	if erro != nil {
		return model.Usuario{}, erro
	}

	usuario.CriadoEm = usuarioBanco.CriadoEm
	usuario.AtualizadoEm = time.Now()

	return u.usuarioRepository.Update(usuario)
}

// Delete -> exclui uma Usuario com o id passado
func (u usuarioService) Delete(id uint64) error {
	return u.usuarioRepository.Delete(id)
}

// FindById -> retorna a Usuario com o id passado
func (u usuarioService) FindById(receitaID uint64, empresaID uint64) (model.Usuario, error) {
	return u.usuarioRepository.FindById(receitaID, empresaID)
}

// FindByUsername -> retorna a Usuario com o username passado
func (u usuarioService) FindByUsername(username string) (model.Usuario, error) {
	return u.usuarioRepository.FindByUsername(username)
}

// GetAll -> retorna todos os Usuarioss cadastrados que contém a descrição desejada
func (u usuarioService) GetAll(usuario model.Usuario, empresaID uint64) ([]model.Usuario, error) {
	return u.usuarioRepository.GetAll(usuario, empresaID)
}
