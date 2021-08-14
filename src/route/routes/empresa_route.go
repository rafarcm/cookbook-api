package routes

import (
	"cookbook/src/controller"
	"cookbook/src/middleware"
	"cookbook/src/repository"
	"cookbook/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetEmpresaRoute : Configura as rotas para Empresas
func GetEmpresaRoute(db *gorm.DB, httpRouter *gin.Engine) *gin.Engine {

	empresaRepository := repository.NewEmpresaRepository(db)
	empresaService := service.NewEmpresaService(empresaRepository)
	empresaController := controller.NewEmpresaController(empresaService)

	empresas := httpRouter.Group("empresas")

	empresas.POST("/", middleware.DBTransactionMiddleware(db), empresaController.AddEmpresa)
	empresas.PUT("/:id", middleware.DBTransactionMiddleware(db), empresaController.UpdateEmpresa)
	empresas.DELETE("/:id", middleware.DBTransactionMiddleware(db), empresaController.DeleteEmpresa)
	empresas.GET("/:id", empresaController.FindEmpresaById)
	empresas.GET("/", empresaController.GetAllEmpresas)
	return httpRouter
}
