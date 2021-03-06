package repository_test

import (
	"cookbook/src/config"
	"cookbook/src/constants"
	"cookbook/src/database"
	"cookbook/src/model"
	"cookbook/src/repository"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func TestIngredienteRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ingrediente Repository Suite")
}

var _ = Describe("IngredienteRepository", func() {
	var (
		tx                   *gorm.DB
		ingredienteRepo      repository.IngredienteRepository
		ingrediente          model.Ingrediente
		ingredientes         []model.Ingrediente
		err                  error
		ingredienteID        uint64
		descricaoIngrediente string
		precoIngrediente     float64 = 10
	)

	BeforeEach(func() {
		config.Carregar("../../.env")
		db, err := database.DBConnection()
		tx = db.Begin()
		Expect(err).To(BeNil())
		ingredienteRepo = repository.NewIngredienteRepository(tx)
	})

	Describe("Save", func() {
		Describe("vai salvar um registro no banco de dados", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente Save"
			})

			It("retorna o registro inserido com sucesso sem erro", func() {
				ingrediente1 := model.Ingrediente{
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}
				ingrediente, err = ingredienteRepo.Save(ingrediente1)
				Expect(err).To(BeNil())
				Expect(ingrediente.ID).NotTo(BeNil())
				Expect(ingrediente.Descricao).To(Equal(descricaoIngrediente))
				Expect(ingrediente.UnidadeMedida).To(Equal(constants.Quilograma))
				Expect(ingrediente.Preco).To(Equal(precoIngrediente))
			})
		})
	})

	Describe("Update", func() {
		Describe("vai atualizar um registro existente no banco de dados", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente Update"
				ingrediente1 := model.Ingrediente{
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}
				err = tx.Create(&ingrediente1).Error
				Expect(err).To(BeNil())
				ingrediente = ingrediente1
			})

			It("retorna o registro atualizado com sucesso sem erro", func() {
				descricaoIngrediente = "Teste Ingrediente Update 2"
				ingrediente.Descricao = descricaoIngrediente
				ingrediente, err = ingredienteRepo.Update(ingrediente)
				Expect(err).To(BeNil())
				Expect(ingrediente.Descricao).To(Equal(descricaoIngrediente))
			})
		})
	})

	Describe("Delete", func() {
		Describe("vai deletar um registro existente no banco de dados", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente Delete"
				ingrediente1 := model.Ingrediente{
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}
				err = tx.Create(&ingrediente1).Error
				Expect(err).To(BeNil())
				ingredienteID = ingrediente1.ID
			})

			It("deleta o registro sem erro", func() {
				err = ingredienteRepo.Delete(ingredienteID)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("FindById", func() {
		Describe("sem registros no banco de dados", func() {
			BeforeEach(func() {
				ingredienteID = 0
			})

			It("n??o retorna nunhum ingrediente", func() {
				ingrediente, _ = ingredienteRepo.FindById(ingredienteID, 1)
				Expect(ingrediente).To(Equal(model.Ingrediente{}))
			})
		})

		Describe("quando existe um registro", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente FindById"
				ingrediente1 := model.Ingrediente{
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}
				err = tx.Create(&ingrediente1).Error
				Expect(err).To(BeNil())
				ingredienteID = ingrediente1.ID
			})

			It("retorna apenas o registro pertencente ao ID", func() {
				ingrediente, err = ingredienteRepo.FindById(ingredienteID, 1)
				Expect(err).To(BeNil())
				Expect(ingrediente.Descricao).To(Equal(descricaoIngrediente))
				Expect(ingrediente.UnidadeMedida).To(Equal(constants.Quilograma))
				Expect(ingrediente.Preco).To(Equal(precoIngrediente))
			})
		})
	})

	Describe("GetAll", func() {
		Describe("sem registros no banco de dados", func() {
			It("n??o retorna nunhum ingrediente", func() {
				ingredientes, err = ingredienteRepo.GetAll("TESTE_SEM_RETORNO", 1, 0)
				Expect(err).To(BeNil())
				Expect(len(ingredientes)).To(Equal(0))
			})
		})

		Describe("quando existe um registro", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente GetAll"
				ingrediente1 := model.Ingrediente{
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}
				err = tx.Create(&ingrediente1).Error
				Expect(err).To(BeNil())
				ingredienteID = ingrediente1.ID
			})

			It("retorna apenas os registros com a descri????o passada", func() {
				ingredientes, err = ingredienteRepo.GetAll(descricaoIngrediente, 1, 0)
				Expect(err).To(BeNil())
				Expect(len(ingredientes)).To(Equal(1))
				Expect(ingredientes[0].Descricao).To(Equal(descricaoIngrediente))
			})
		})
	})

	Describe("WithTrx", func() {
		Describe("vai retornar um repository com transa????o", func() {
			It("retorna o service com transacao", func() {
				ingredienteRepo.WithTrx(tx)
			})
		})

		Describe("vai retornar um repository sem transa????o", func() {
			It("retorna o service com transacao", func() {
				ingredienteRepo.WithTrx(nil)
			})
		})
	})

	AfterEach(func() {
		tx.Rollback()
	})
})
