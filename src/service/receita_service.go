package service

import (
	"cookbook/src/model"
	"cookbook/src/repository"
	"time"

	"gorm.io/gorm"
)

type receitaService struct {
	receitaRepository repository.ReceitaRepository
}

// ReceitaService : representa o contrato de Receita service
type ReceitaService interface {
	WithTrx(*gorm.DB) ReceitaService
	Save(model.Receita) (model.Receita, error)
	Update(model.Receita) (model.Receita, error)
	Delete(uint64) error
	FindById(uint64) (model.Receita, error)
}

// NewReceitaService -> retorna um novo Receita service
func NewReceitaService(r repository.ReceitaRepository) ReceitaService {
	return receitaService{
		receitaRepository: r,
	}
}

// WithTrx : habilita repositório com transação
func (r receitaService) WithTrx(trxHandle *gorm.DB) ReceitaService {
	r.receitaRepository = r.receitaRepository.WithTrx(trxHandle)
	return r
}

// Save -> salva uma nova receita e a retorna
func (r receitaService) Save(receita model.Receita) (model.Receita, error) {
	receita.CriadoEm = time.Now()
	receita.AtualizadoEm = receita.CriadoEm

	return r.receitaRepository.Save(receita)
}

// Update -> atualiza a receita e a retorna
func (r receitaService) Update(receita model.Receita) (model.Receita, error) {
	receitaBanco, erro := r.receitaRepository.FindById(receita.ID)
	if erro != nil {
		return model.Receita{}, erro
	}

	receita.CriadoEm = receitaBanco.CriadoEm
	receita.AtualizadoEm = time.Now()

	return r.receitaRepository.Update(receita)
}

// Delete -> exclui uma receita com o id passado
func (r receitaService) Delete(id uint64) error {
	return r.receitaRepository.Delete(id)
}

// FindById -> retorna a receita com o id passado
func (r receitaService) FindById(id uint64) (model.Receita, error) {
	return r.receitaRepository.FindById(id)
}
