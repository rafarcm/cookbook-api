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

// IngredienteService : Representa o contrato de ingredientes service
type IngredienteService interface {
	WithTrx(*gorm.DB) IngredienteService
	Save(model.Ingrediente) (model.Ingrediente, error)
	Update(uint64, model.Ingrediente) (model.Ingrediente, error)
	Delete(uint64) error
	FindById(uint64) (model.Ingrediente, error)
	GetAll(string) ([]model.Ingrediente, error)
}

// NewIngredienteService -> retorna um novo ingrediente service
func NewIngredienteService(r repository.IngredienteRepository) IngredienteService {
	return ingredienteService{
		ingredienteRepository: r,
	}
}

// WithTrx : Habilita repositório com transação
func (i ingredienteService) WithTrx(trxHandle *gorm.DB) IngredienteService {
	i.ingredienteRepository = i.ingredienteRepository.WithTrx(trxHandle)
	return i
}

// Save -> Salva um novo ingrediente e o retorna
func (i ingredienteService) Save(ingrediente model.Ingrediente) (model.Ingrediente, error) {
	ingrediente.CriadoEm = time.Now()
	ingrediente.AtualizadoEm = ingrediente.CriadoEm

	return i.ingredienteRepository.Save(ingrediente)
}

// Update -> Atualiza a descrição e unidade de medidao do ingediente
func (i ingredienteService) Update(id uint64, ingrediente model.Ingrediente) (model.Ingrediente, error) {
	ingredienteBanco, erro := i.ingredienteRepository.FindById(id)

	if erro != nil {
		return model.Ingrediente{}, erro
	}

	ingredienteBanco.AtualizadoEm = time.Now()
	ingredienteBanco.Descricao = ingrediente.Descricao
	ingredienteBanco.UnidadeMedida = ingrediente.UnidadeMedida

	return i.ingredienteRepository.Update(ingredienteBanco)
}

// Delete -> Exclui um ingediente com o id passado
func (i ingredienteService) Delete(id uint64) error {
	return i.ingredienteRepository.Delete(id)
}

// FindById -> Retorna o ingediente com o id passado
func (i ingredienteService) FindById(id uint64) (model.Ingrediente, error) {
	return i.ingredienteRepository.FindById(id)
}

// GetAll -> Retorna todos os ingredientes cadastrados que contém a descrição desejada
func (i ingredienteService) GetAll(descricao string) ([]model.Ingrediente, error) {
	return i.ingredienteRepository.GetAll(descricao)
}
