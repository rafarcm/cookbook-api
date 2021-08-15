package service_test

import (
	"cookbook/src/constants"
	mock "cookbook/src/mock/repository"
	"cookbook/src/model"
	"cookbook/src/repository"
	"cookbook/src/service"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func TestIngredienteService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ingrediente Service Suite")
}

var _ = Describe("OrderService", func() {
	var (
		ingredienteRepo        repository.IngredienteRepository
		ingredienteRepoWithTrx repository.IngredienteRepository
		ingredienteService     service.IngredienteService
		ingrediente            model.Ingrediente
		ingredientes           []model.Ingrediente
		err                    error
		ctrl                   *gomock.Controller
		db                     *gorm.DB
		descricaoIngrediente   string
		precoIngrediente       float64 = 10
		ingredienteID          uint64  = 10
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	JustBeforeEach(func() {
		ingredienteServiceImpl := service.NewIngredienteService(ingredienteRepo)
		ingredienteService = ingredienteServiceImpl
	})

	Describe("Save", func() {
		Describe("vai salvar um registro no banco de dados", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente Save"
				ingrediente = model.Ingrediente{
					ID:            ingredienteID,
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}

				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().
					Save(gomock.Any()).
					Return(ingrediente, error(nil))
				ingredienteRepo = ingredienteRepoMock
			})

			It("retorna o registro inserido com sucesso sem erro", func() {
				ingrediente, err = ingredienteService.Save(ingrediente)
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
				descricaoIngrediente = "Teste Ingrediente Após Update"
				ingrediente = model.Ingrediente{
					ID:            ingredienteID,
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}

				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().
					FindById(gomock.Eq(ingredienteID)).
					Return(ingrediente, error(nil))
				ingredienteRepoMock.EXPECT().
					Update(gomock.Any()).
					Return(ingrediente, error(nil))
				ingredienteRepo = ingredienteRepoMock
			})

			It("retorna o registro atualizado com sucesso sem erro", func() {
				ingrediente.ID = ingredienteID
				ingrediente, err = ingredienteService.Update(ingrediente)
				Expect(err).To(BeNil())
				Expect(ingrediente.ID).NotTo(BeNil())
				Expect(ingrediente.Descricao).To(Equal(descricaoIngrediente))
				Expect(ingrediente.UnidadeMedida).To(Equal(constants.Quilograma))
				Expect(ingrediente.Preco).To(Equal(precoIngrediente))
			})
		})

		Describe("vai retornar erro ao tentar buscar o ingrediente com o id passado", func() {
			BeforeEach(func() {
				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().
					FindById(gomock.Eq(ingredienteID)).
					Return(model.Ingrediente{}, gorm.ErrRecordNotFound)
				ingredienteRepo = ingredienteRepoMock
			})

			It("retorna erro de registro não encontrado", func() {
				ingrediente.ID = ingredienteID
				ingrediente, err = ingredienteService.Update(ingrediente)
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
				Expect(ingrediente).To(Equal(model.Ingrediente{}))
			})
		})
	})

	Describe("Delete", func() {
		Describe("vai deletar um registro existente no banco de dados", func() {
			BeforeEach(func() {
				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)

				ingredienteRepoMock.EXPECT().
					Delete(gomock.Eq(ingredienteID)).
					Return(error(nil))
				ingredienteRepo = ingredienteRepoMock
			})

			It("deleta o registro sem erro", func() {
				err = ingredienteService.Delete(ingredienteID)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("FindById", func() {
		Describe("sem registros no banco de dados", func() {
			BeforeEach(func() {
				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().FindById(gomock.Eq(ingredienteID))
				ingredienteRepo = ingredienteRepoMock
			})

			It("não retorna nunhum ingrediente", func() {
				ingrediente, _ = ingredienteService.FindById(ingredienteID, 1)
				Expect(ingrediente).To(Equal(model.Ingrediente{}))
			})
		})

		Describe("quando temos registro a ser retornado", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente FindById"
				ingrediente1 := model.Ingrediente{
					ID:            ingredienteID,
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}

				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().
					FindById(gomock.Eq(ingredienteID)).
					Return(ingrediente1, error(nil))
				ingredienteRepo = ingredienteRepoMock
			})

			It("retorna o ingrediente", func() {
				ingrediente, _ = ingredienteService.FindById(ingredienteID, 1)
				Expect(ingrediente.ID).To(Equal(ingredienteID))
				Expect(ingrediente.Descricao).To(Equal(descricaoIngrediente))
				Expect(ingrediente.UnidadeMedida).To(Equal(constants.Quilograma))
				Expect(ingrediente.Preco).To(Equal(precoIngrediente))
			})
		})
	})

	Describe("GetAll", func() {
		Describe("sem registros no banco de dados", func() {
			BeforeEach(func() {
				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().GetAll(gomock.Eq(""))
				ingredienteRepo = ingredienteRepoMock
			})

			It("não retorna nunhum ingrediente", func() {
				ingredientes, err = ingredienteService.GetAll("", 1)
				Expect(err).To(BeNil())
				Expect(len(ingredientes)).To(Equal(0))
			})
		})

		Describe("quando existe um registro", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente GetAll"
				ingrediente1 := model.Ingrediente{
					ID:            1,
					Descricao:     "Teste Ingrediente GetAll 1",
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}

				ingrediente2 := model.Ingrediente{
					ID:            2,
					Descricao:     "Teste Ingrediente GetAll 2",
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}

				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().
					GetAll(gomock.Eq(descricaoIngrediente)).
					Return([]model.Ingrediente{ingrediente1, ingrediente2}, error(nil))
				ingredienteRepo = ingredienteRepoMock
			})

			It("retorna apenas os registros com a descrição passada", func() {
				ingredientes, err = ingredienteService.GetAll(descricaoIngrediente, 1)
				Expect(err).To(BeNil())
				Expect(len(ingredientes)).To(Equal(2))
			})
		})
	})

	Describe("WithTrx", func() {
		Describe("vai retornar um service com transação", func() {
			BeforeEach(func() {
				ingredienteRepoMock := mock.NewMockIngredienteRepository(ctrl)

				ingredienteRepoMock.EXPECT().
					WithTrx(gomock.Any()).
					Return(ingredienteRepoWithTrx)
				ingredienteRepo = ingredienteRepoMock
			})

			It("retorna o service com transacao", func() {
				ingredienteService.WithTrx(db)
			})
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
