package service

import (
	"cookbook/src/model"
	"cookbook/src/repository"
	"time"

	"gorm.io/gorm"
)

type empresaService struct {
	empresaRepository repository.EmpresaRepository
}

// EmpresaService : representa o contrato de Empresa service
type EmpresaService interface {
	WithTrx(*gorm.DB) EmpresaService
	Save(model.Empresa) (model.Empresa, error)
	Update(model.Empresa) (model.Empresa, error)
	Delete(uint64) error
	FindById(uint64) (model.Empresa, error)
	GetAll(string) ([]model.Empresa, error)
}

// NewEmpresaService -> retorna um novo Empresa service
func NewEmpresaService(u repository.EmpresaRepository) EmpresaService {
	return empresaService{
		empresaRepository: u,
	}
}

// WithTrx : habilita repositório com transação
func (e empresaService) WithTrx(trxHandle *gorm.DB) EmpresaService {
	e.empresaRepository = e.empresaRepository.WithTrx(trxHandle)
	return e
}

// Save -> salva uma nova Empresa e a retorna
func (e empresaService) Save(empresa model.Empresa) (model.Empresa, error) {
	empresa.CriadoEm = time.Now()
	empresa.AtualizadoEm = empresa.CriadoEm

	for i := range empresa.Usuarios {
		empresa.Usuarios[i].CriadoEm = empresa.CriadoEm
		empresa.Usuarios[i].AtualizadoEm = empresa.AtualizadoEm
	}

	return e.empresaRepository.Save(empresa)
}

// Update -> atualiza a Empresa e a retorna
func (e empresaService) Update(empresa model.Empresa) (model.Empresa, error) {
	EmpresaBanco, erro := e.empresaRepository.FindById(empresa.ID)
	if erro != nil {
		return model.Empresa{}, erro
	}

	empresa.CriadoEm = EmpresaBanco.CriadoEm
	empresa.AtualizadoEm = time.Now()

	return e.empresaRepository.Update(empresa)
}

// Delete -> exclui uma Empresa com o id passado
func (e empresaService) Delete(id uint64) error {
	return e.empresaRepository.Delete(id)
}

// FindById -> retorna a Empresa com o id passado
func (e empresaService) FindById(id uint64) (model.Empresa, error) {
	return e.empresaRepository.FindById(id)
}

// GetAll -> retorna todos os Empresass cadastrados que contém a descrição desejada
func (e empresaService) GetAll(descricao string) ([]model.Empresa, error) {
	return e.empresaRepository.GetAll(descricao)
}
