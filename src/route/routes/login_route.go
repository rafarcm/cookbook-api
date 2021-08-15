package routes

import (
	"cookbook/src/controller"
	"cookbook/src/repository"
	"cookbook/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetLoginRoute : Configura as rotas para Login
func GetLoginRoute(db *gorm.DB, httpRouter *gin.Engine) *gin.Engine {

	usuarioRepository := repository.NewUsuarioRepository(db)
	usuarioService := service.NewUsuarioService(usuarioRepository)
	loginController := controller.NewLoginController(usuarioService)

	receitas := httpRouter.Group("login")

	receitas.POST("/", loginController.Login)
	return httpRouter
}
