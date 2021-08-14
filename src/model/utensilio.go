package model

import "time"

type Utensilio struct {
	ID           uint64    `json:"id,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Descricao    string    `json:"descricao,omitempty" binding:"required" gorm:"not null"`
	CriadoEm     time.Time `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time `json:"atualizadoEm" gorm:"not null"`
}
