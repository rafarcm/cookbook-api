package service

import (
	"cookbook/src/constants"
	"cookbook/src/model"
	"cookbook/src/repository"
	"errors"
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
	Update(model.Receita, uint64) (model.Receita, error)
	Delete(uint64, uint64) error
	FindById(uint64, uint64) (model.Receita, error)
	GetAll(string, constants.Categoria, uint64, uint64) ([]model.Receita, error)
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
func (r receitaService) Update(receita model.Receita, empresaID uint64) (model.Receita, error) {
	receitaBanco, erro := r.receitaRepository.FindById(receita.ID, empresaID)
	if erro != nil {
		return model.Receita{}, erro
	}

	if receitaBanco.ID == 0 {
		return model.Receita{}, errors.New("não é possível alterar esta receita")
	}

	receita.CriadoEm = receitaBanco.CriadoEm
	receita.AtualizadoEm = time.Now()

	return r.receitaRepository.Update(receita)
}

// Delete -> exclui uma receita com o id passado
func (r receitaService) Delete(id uint64, empresaID uint64) error {
	receitaBanco, erro := r.receitaRepository.FindById(id, empresaID)
	if erro != nil {
		return erro
	}

	if receitaBanco.ID == 0 {
		return errors.New("não é possível deletar esta receita")
	}
	return r.receitaRepository.Delete(id)
}

// FindById -> retorna a receita com o id passado
func (r receitaService) FindById(receitaID uint64, empresaID uint64) (model.Receita, error) {
	return r.receitaRepository.FindById(receitaID, empresaID)
}

// GetAll -> retorna todas as receitas cadastradas de acordo com os parâmetros passados
func (r receitaService) GetAll(descricao string, categoria constants.Categoria, empresaID uint64, usuarioID uint64) (receitas []model.Receita, erro error) {
	return r.receitaRepository.GetAll(descricao, categoria, empresaID, usuarioID)
}
