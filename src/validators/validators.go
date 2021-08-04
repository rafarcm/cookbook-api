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
