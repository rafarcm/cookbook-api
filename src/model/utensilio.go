package model

import "time"

// Utensilio ... utensilio Database Model
type Utensilio struct {
	ID           uint64    `json:"utensilioId,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Descricao    string    `json:"descricao,omitempty" binding:"required" gorm:"not null"`
	EmpresaID    uint64    `json:"empresaId,omitempty" binding:"required" gorm:"not null"`
	CriadoEm     time.Time `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time `json:"atualizadoEm" gorm:"not null"`
}
