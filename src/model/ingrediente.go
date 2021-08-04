package model

import "time"

// Ingrediente ... Ingrediente Database Model
type Ingrediente struct {
	ID            uint64             `json:"id,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Descricao     string             `json:"descricao,omitempty" binding:"required" gorm:"not null"`
	UnidadeMedida string             `json:"unidadeMedida,omitempty" binding:"required,unidademedidavalidation" gorm:"not null"`
	Precos        []PrecoIngrediente `json:"precos,omitempty" binding:"precosvalidation" gorm:"foreignkey:IngredienteID;constraint:OnDelete:CASCADE;"`
	CriadoEm      time.Time          `json:"criadoEm" gorm:"not null"`
	AtualizadoEm  time.Time          `json:"atualizadoEm" gorm:"not null"`
}

type Ingredientes []*Ingrediente
