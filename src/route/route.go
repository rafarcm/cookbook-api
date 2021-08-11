package route

import (
	"cookbook/src/route/routes"
	"cookbook/src/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

//SetupRoutes : Configura as rotas da API
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unidademedidavalidation", validators.UnidadeMedidaValidation)
	}

	httpRouter = routes.GetIngredienteRoute(db, httpRouter)
	httpRouter = routes.GetReceitaRoute(db, httpRouter)
	httpRouter.Run()
}
