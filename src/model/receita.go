package model

import (
	"cookbook/src/constants"
	"time"
)

// Receita ... Receita Database Model
type Receita struct {
	ID            uint64               `json:"id,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Descricao     string               `json:"descricao,omitempty" binding:"required" gorm:"not null"`
	ModoPreparo   string               `json:"modoPreparo,omitempty" binding:"required" gorm:"not null"`
	Rendimento    uint64               `json:"rendimento,omitempty"`
	TempoPreparo  uint64               `json:"tempoPreparo,omitempty"`
	Conservacao   uint16               `json:"conservacao,omitempty"`
	Categoria     constants.Categoria  `json:"categoria,omitempty" binding:"required,categoriavalidation" gorm:"not null"`
	Preco         float64              `json:"preco,omitempty" binding:"required" gorm:"not null"`
	PrecoSugerido float64              `json:"precoSugerido,omitempty" gorm:"not null"`
	Ingredientes  []IngredienteReceita `json:"ingredientes,omitempty" binding:"required" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Utensilios    []Utensilio          `json:"utensilios,omitempty" binding:"required" gorm:"many2many:receita_utensilios;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EmpresaID     uint64               `json:"empresa,omitempty" binding:"required" gorm:"not null"`
	CriadoEm      time.Time            `json:"criadoEm" gorm:"not null"`
	AtualizadoEm  time.Time            `json:"atualizadoEm" gorm:"not null"`
}
