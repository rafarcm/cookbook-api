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
	FindById(uint64) (model.Utensilio, error)
	GetAll(string) ([]model.Utensilio, error)
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
func (u utensilioService) Save(Utensilio model.Utensilio) (model.Utensilio, error) {
	Utensilio.CriadoEm = time.Now()
	Utensilio.AtualizadoEm = Utensilio.CriadoEm

	return u.utensilioRepository.Save(Utensilio)
}

// Update -> atualiza a Utensilio e a retorna
func (u utensilioService) Update(Utensilio model.Utensilio) (model.Utensilio, error) {
	UtensilioBanco, erro := u.utensilioRepository.FindById(Utensilio.ID)
	if erro != nil {
		return model.Utensilio{}, erro
	}

	Utensilio.CriadoEm = UtensilioBanco.CriadoEm
	Utensilio.AtualizadoEm = time.Now()

	return u.utensilioRepository.Update(Utensilio)
}

// Delete -> exclui uma Utensilio com o id passado
func (u utensilioService) Delete(id uint64) error {
	return u.utensilioRepository.Delete(id)
}

// FindById -> retorna a Utensilio com o id passado
func (u utensilioService) FindById(id uint64) (model.Utensilio, error) {
	return u.utensilioRepository.FindById(id)
}

// GetAll -> retorna todos os utensilioss cadastrados que contém a descrição desejada
func (u utensilioService) GetAll(descricao string) ([]model.Utensilio, error) {
	return u.utensilioRepository.GetAll(descricao)
}
