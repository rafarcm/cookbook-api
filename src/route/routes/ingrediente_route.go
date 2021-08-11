package routes

import (
	"cookbook/src/controller"
	"cookbook/src/middleware"
	"cookbook/src/repository"
	"cookbook/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetIngredienteRoute : Configura as rotas para Ingredientes
func GetIngredienteRoute(db *gorm.DB, httpRouter *gin.Engine) *gin.Engine {

	ingredienteRepository := repository.NewIngredienteRepository(db)
	ingredienteService := service.NewIngredienteService(ingredienteRepository)
	ingredienteController := controller.NewIngredienteController(ingredienteService)

	ingredientes := httpRouter.Group("ingredientes")

	ingredientes.POST("/", middleware.DBTransactionMiddleware(db), ingredienteController.AddIngrediente)
	ingredientes.PUT("/:id", middleware.DBTransactionMiddleware(db), ingredienteController.UpdateIngrediente)
	ingredientes.DELETE("/:id", middleware.DBTransactionMiddleware(db), ingredienteController.DeleteIngrediente)
	ingredientes.GET("/:id", ingredienteController.FindIngredienteById)
	ingredientes.GET("/", ingredienteController.GetAllIngredientes)

	return httpRouter
}
