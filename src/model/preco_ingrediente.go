package model

import (
	"time"
)

type PrecoIngrediente struct {
	ID            uint64    `json:"id,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Preco         float64   `json:"preco,omitempty" binding:"required" gorm:"not null"`
	CriadoEm      time.Time `json:"criadoEm,omitempty" gorm:"not null"`
	IngredienteID uint64    `json:"ingredienteId,omitempty" gorm:"not null"`
}
