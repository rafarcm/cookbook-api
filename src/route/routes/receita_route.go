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

	receitas.POST("/", middleware.DBTransactionMiddleware(db), receitaController.AddReceita)
	receitas.PUT("/:id", middleware.DBTransactionMiddleware(db), receitaController.UpdateReceita)
	receitas.DELETE("/:id", middleware.DBTransactionMiddleware(db), receitaController.DeleteReceita)
	receitas.GET("/:id", receitaController.FindReceitaById)
	receitas.GET("/", receitaController.GetAllReceitas)
	return httpRouter
}
