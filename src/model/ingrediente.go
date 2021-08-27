package model

import "time"

// Ingrediente ... ingrediente Database Model
type Ingrediente struct {
	ID            uint64               `json:"ingredienteId,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Descricao     string               `json:"descricao,omitempty" binding:"required" gorm:"not null"`
	UnidadeMedida string               `json:"unidadeMedida,omitempty" binding:"required,unidademedidavalidation" gorm:"not null"`
	Preco         float64              `json:"preco,omitempty" gorm:"not null"`
	Receitas      []IngredienteReceita `json:"receitas,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UsuarioID     uint64               `json:"usuarioId,omitempty" gorm:"not null"`
	CriadoEm      time.Time            `json:"criadoEm" gorm:"not null"`
	AtualizadoEm  time.Time            `json:"atualizadoEm" gorm:"not null"`
}
