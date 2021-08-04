package validators

import (
	"cookbook/src/constants"
	"cookbook/src/model"

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

var PrecosValidation validator.Func = func(fl validator.FieldLevel) bool {
	precos, ok := fl.Field().Interface().([]model.PrecoIngrediente)
	if ok {
		if len(precos) <= 0 {
			return false
		}

		for idx := range precos {
			if precos[idx].Preco <= 0 {
				return false
			}
		}
	}
	return true
}
