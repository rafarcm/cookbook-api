package validators

import (
	"cookbook/src/constants"

	"github.com/go-playground/validator/v10"
)

var UnidadeMedidaValidation validator.Func = func(fl validator.FieldLevel) bool {
	str, ok := fl.Field().Interface().(string)
	if ok {
		for _, valor := range constants.UnidadesMedida {
			if str == valor {
				return true
			}
		}
	}
	return false
}

var CategoriaValidation validator.Func = func(fl validator.FieldLevel) bool {
	code, ok := fl.Field().Interface().(constants.Categoria)
	defer categoriaInvalida()
	if ok {
		_ = constants.Categoria(code).String()
	}
	return ok
}

func categoriaInvalida() {
	if r := recover(); r != nil {
		return
	}
}
