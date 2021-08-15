package service

import (
	"cookbook/src/model"
	"cookbook/src/repository"
	"time"

	"gorm.io/gorm"
)

type utensilioService struct {
	utensilioRepository repository.UtensilioRepository
}

// UtensilioService : representa o contrato de Utensilio service
type UtensilioService interface {
	WithTrx(*gorm.DB) UtensilioService
	Save(model.Utensilio) (model.Utensilio, error)
	Update(model.Utensilio) (model.Utensilio, error)
	Delete(uint64) error
	FindById(uint64, uint64) (model.Utensilio, error)
	GetAll(string, uint64) ([]model.Utensilio, error)
}

// NewUtensilioService -> retorna um novo Utensilio service
func NewUtensilioService(u repository.UtensilioRepository) UtensilioService {
	return utensilioService{
		utensilioRepository: u,
	}
}

// WithTrx : habilita repositório com transação
func (u utensilioService) WithTrx(trxHandle *gorm.DB) UtensilioService {
	u.utensilioRepository = u.utensilioRepository.WithTrx(trxHandle)
	return u
}

// Save -> salva uma nova Utensilio e a retorna
func (u utensilioService) Save(utensilio model.Utensilio) (model.Utensilio, error) {
	utensilio.CriadoEm = time.Now()
	utensilio.AtualizadoEm = utensilio.CriadoEm

	return u.utensilioRepository.Save(utensilio)
}

// Update -> atualiza a Utensilio e a retorna
func (u utensilioService) Update(utensilio model.Utensilio) (model.Utensilio, error) {
	utensilioBanco, erro := u.utensilioRepository.FindById(utensilio.ID, utensilio.EmpresaID)
	if erro != nil {
		return model.Utensilio{}, erro
	}

	utensilio.CriadoEm = utensilioBanco.CriadoEm
	utensilio.AtualizadoEm = time.Now()

	return u.utensilioRepository.Update(utensilio)
}

// Delete -> exclui uma Utensilio com o id passado
func (u utensilioService) Delete(id uint64) error {
	return u.utensilioRepository.Delete(id)
}

// FindById -> retorna a Utensilio com o id passado
func (u utensilioService) FindById(utensilioID uint64, empresaID uint64) (model.Utensilio, error) {
	return u.utensilioRepository.FindById(utensilioID, empresaID)
}

// GetAll -> retorna todos os utensilioss cadastrados que contém a descrição desejada
func (u utensilioService) GetAll(descricao string, empresaID uint64) ([]model.Utensilio, error) {
	return u.utensilioRepository.GetAll(descricao, empresaID)
}
