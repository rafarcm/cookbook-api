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
	UtensilioService := service.NewUtensilioService(utensilioRepository)
	utensilioController := controller.NewUtensilioController(UtensilioService)

	utensilios := httpRouter.Group("utensilios")

	utensilios.POST("/", middleware.DBTransactionMiddleware(db), utensilioController.AddUtensilio)
	utensilios.PUT("/:id", middleware.DBTransactionMiddleware(db), utensilioController.UpdateUtensilio)
	utensilios.DELETE("/:id", middleware.DBTransactionMiddleware(db), utensilioController.DeleteUtensilio)
	utensilios.GET("/:id", utensilioController.FindUtensilioById)
	utensilios.GET("/", utensilioController.GetAllUtensilios)

	return httpRouter
}
