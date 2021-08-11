package model

import "time"

// Receita ... Receita Database Model
type Receita struct {
	ID            uint64               `json:"id,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Descricao     string               `json:"descricao,omitempty" binding:"required" gorm:"not null"`
	ModoPreparo   string               `json:"modoPreparo,omitempty" binding:"required" gorm:"not null"`
	Rendimento    uint64               `json:"rendimento,omitempty"`
	TempoPreparo  uint64               `json:"tempoPreparo,omitempty"`
	Preco         float64              `json:"preco,omitempty" gorm:"not null"`
	PrecoSugerido float64              `json:"precoSugerido,omitempty" gorm:"not null"`
	Ingredientes  []IngredienteReceita `json:"ingredientes,omitempty" binding:"required" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CriadoEm      time.Time            `json:"criadoEm" gorm:"not null"`
	AtualizadoEm  time.Time            `json:"atualizadoEm" gorm:"not null"`
}
