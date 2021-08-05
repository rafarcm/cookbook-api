package service_test

import (
	"cookbook/src/constants"
	"cookbook/src/mock_repository"
	"cookbook/src/model"
	"cookbook/src/repository"
	"cookbook/src/service"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIngredienteService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ingrediente Service Suite")
}

var _ = Describe("OrderService", func() {
	var (
		ingredienteRepo      repository.IngredienteRepository
		ingredienteService   service.IngredienteService
		ingrediente          model.Ingrediente
		ctrl                 *gomock.Controller
		descricaoIngrediente string
		precoIngrediente     float64 = 10
		userID               uint64  = 5
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	JustBeforeEach(func() {
		ingredienteServiceImpl := service.NewIngredienteService(ingredienteRepo)
		ingredienteService = ingredienteServiceImpl
	})

	Describe("FindById", func() {
		Describe("sem registros no banco de dados", func() {
			BeforeEach(func() {
				ingredienteRepoMock := mock_repository.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().FindById(gomock.Eq(userID))
				ingredienteRepo = ingredienteRepoMock
			})

			It("não retorna nunhum ingrediente", func() {
				ingrediente, _ = ingredienteService.FindById(userID)
				Expect(ingrediente.Descricao).To(Equal(""))
				Expect(ingrediente.UnidadeMedida).To(Equal(""))
			})
		})

		Describe("quando temos registro a ser retornado", func() {
			BeforeEach(func() {
				descricaoIngrediente = "Teste Ingrediente FindById"
				userID = 10
				precoIngrediente = 10
				ingrediente1 := model.Ingrediente{
					ID:            userID,
					Descricao:     descricaoIngrediente,
					UnidadeMedida: constants.Quilograma,
					Preco:         precoIngrediente,
					CriadoEm:      time.Now(),
					AtualizadoEm:  time.Now(),
				}

				ingredienteRepoMock := mock_repository.NewMockIngredienteRepository(ctrl)
				ingredienteRepoMock.EXPECT().
					FindById(gomock.Eq(userID)).
					Return(ingrediente1, error(nil))
				ingredienteRepo = ingredienteRepoMock
			})

			It("retorna o ingrediente", func() {
				ingrediente, _ = ingredienteService.FindById(userID)
				Expect(ingrediente.ID).To(Equal(userID))
				Expect(ingrediente.Descricao).To(Equal(descricaoIngrediente))
				Expect(ingrediente.UnidadeMedida).To(Equal(constants.Quilograma))
				Expect(ingrediente.Preco).To(Equal(precoIngrediente))
			})
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
