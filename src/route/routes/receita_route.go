package routes

import (
	"cookbook/src/controller"
	"cookbook/src/middleware"
	"cookbook/src/repository"
	"cookbook/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetReceitaRoute : Configura as rotas para Receitas
func GetReceitaRoute(db *gorm.DB, httpRouter *gin.Engine) *gin.Engine {

	receitaRepository := repository.NewReceitaRepository(db)
	receitaService := service.NewReceitaService(receitaRepository)
	receitaController := controller.NewReceitaController(receitaService)

	receitas := httpRouter.Group("receitas")

	receitas.POST("/", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), receitaController.AddReceita)
	receitas.PUT("/:id", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), receitaController.UpdateReceita)
	receitas.DELETE("/:id", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), receitaController.DeleteReceita)
	receitas.GET("/:id", middleware.Autenticar(), receitaController.FindReceitaById)
	receitas.GET("/", middleware.Autenticar(), receitaController.GetAllReceitas)
	return httpRouter
}
