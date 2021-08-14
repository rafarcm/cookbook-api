package model

import "time"

// Empresa ... empresa Database Model
type Empresa struct {
	ID           uint64    `json:"empresaId,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Nome         string    `json:"nome,omitempty" binding:"required" gorm:"not null"`
	Usuarios     []Usuario `json:"usuarios,omitempty" gorm:"foreignKey:EmpresaID"`
	CriadoEm     time.Time `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time `json:"atualizadoEm" gorm:"not null"`
}
