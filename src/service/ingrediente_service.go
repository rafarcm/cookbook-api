package service

import (
	"cookbook/src/model"
	"cookbook/src/repository"
	"time"

	"gorm.io/gorm"
)

type ingredienteService struct {
	ingredienteRepository repository.IngredienteRepository
}

// IngredienteService : representa o contrato de ingredientes service
type IngredienteService interface {
	WithTrx(*gorm.DB) IngredienteService
	Save(model.Ingrediente) (model.Ingrediente, error)
	Update(model.Ingrediente) (model.Ingrediente, error)
	Delete(uint64) error
	FindById(uint64, uint64) (model.Ingrediente, error)
	GetAll(string, uint64) ([]model.Ingrediente, error)
}

// NewIngredienteService -> retorna um novo ingrediente service
func NewIngredienteService(r repository.IngredienteRepository) IngredienteService {
	return ingredienteService{
		ingredienteRepository: r,
	}
}

// WithTrx : habilita repositório com transação
func (i ingredienteService) WithTrx(trxHandle *gorm.DB) IngredienteService {
	i.ingredienteRepository = i.ingredienteRepository.WithTrx(trxHandle)
	return i
}

// Save -> salva um novo ingrediente e o retorna
func (i ingredienteService) Save(ingrediente model.Ingrediente) (model.Ingrediente, error) {
	ingrediente.CriadoEm = time.Now()
	ingrediente.AtualizadoEm = ingrediente.CriadoEm

	return i.ingredienteRepository.Save(ingrediente)
}

// Update -> atualiza a descrição e unidade de medida do ingrediente e o retorna
func (i ingredienteService) Update(ingrediente model.Ingrediente) (model.Ingrediente, error) {
	ingredienteBanco, erro := i.ingredienteRepository.FindById(ingrediente.ID, ingrediente.EmpresaID)
	if erro != nil {
		return model.Ingrediente{}, erro
	}

	ingrediente.CriadoEm = ingredienteBanco.CriadoEm
	ingrediente.AtualizadoEm = time.Now()

	return i.ingredienteRepository.Update(ingredienteBanco)
}

// Delete -> exclui um ingrediente com o id passado
func (i ingredienteService) Delete(id uint64) error {
	return i.ingredienteRepository.Delete(id)
}

// FindById -> retorna o ingrediente com o id passado
func (i ingredienteService) FindById(ingredienteID uint64, empresaID uint64) (model.Ingrediente, error) {
	return i.ingredienteRepository.FindById(ingredienteID, empresaID)
}

// GetAll -> retorna todos os ingredientes cadastrados que contém a descrição desejada
func (i ingredienteService) GetAll(descricao string, empresaID uint64) ([]model.Ingrediente, error) {
	return i.ingredienteRepository.GetAll(descricao, empresaID)
}
