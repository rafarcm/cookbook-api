package routes

import (
	"cookbook/src/controller"
	"cookbook/src/middleware"
	"cookbook/src/repository"
	"cookbook/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUtensiliooute : Configura as rotas para Utensilios
func GetUtensilioRoute(db *gorm.DB, httpRouter *gin.Engine) *gin.Engine {

	utensilioRepository := repository.NewUtensilioRepository(db)
	utensilioService := service.NewUtensilioService(utensilioRepository)
	utensilioController := controller.NewUtensilioController(utensilioService)

	utensilios := httpRouter.Group("utensilios")

	utensilios.POST("/", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), utensilioController.AddUtensilio)
	utensilios.PUT("/:id", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), utensilioController.UpdateUtensilio)
	utensilios.DELETE("/:id", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), utensilioController.DeleteUtensilio)
	utensilios.GET("/:id", middleware.Autenticar(), utensilioController.FindUtensilioById)
	utensilios.GET("/", middleware.Autenticar(), utensilioController.GetAllUtensilios)

	return httpRouter
}
