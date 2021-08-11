package model

// IngredienteReceita ... IngredienteReceita Database Model
type IngredienteReceita struct {
	IngredienteID uint64  `json:"ingredienteID,omitempty" gorm:"primaryKey"`
	ReceitaID     uint64  `json:"receitaID,omitempty" gorm:"primaryKey"`
	UnidadeMedida string  `json:"unidadeMedida,omitempty" binding:"required,unidademedidavalidation" gorm:"not null"`
	Quantidade    float64 `json:"quantidade,omitempty" binding:"required" gorm:"not null"`
}
