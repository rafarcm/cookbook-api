package routes

import (
	"cookbook/src/controller"
	"cookbook/src/middleware"
	"cookbook/src/repository"
	"cookbook/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsuarioRoute : Configura as rotas para Usuarios
func GetUsuarioRoute(db *gorm.DB, httpRouter *gin.Engine) *gin.Engine {

	usuarioRepository := repository.NewUsuarioRepository(db)
	usuarioService := service.NewUsuarioService(usuarioRepository)
	usuarioController := controller.NewUsuarioController(usuarioService)

	usuarios := httpRouter.Group("usuarios")

	usuarios.POST("/", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), usuarioController.AddUsuario)
	usuarios.PUT("/:id", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), usuarioController.UpdateUsuario)
	usuarios.DELETE("/:id", middleware.Autenticar(), middleware.DBTransactionMiddleware(db), usuarioController.DeleteUsuario)
	usuarios.GET("/:id", middleware.Autenticar(), usuarioController.FindUsuarioById)
	usuarios.GET("/", middleware.Autenticar(), usuarioController.GetAllUsuarios)
	return httpRouter
}
