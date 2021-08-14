package model

import "time"

type Empresa struct {
	ID           uint64    `json:"id,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Nome         string    `json:"nome,omitempty" binding:"required" gorm:"not null"`
	Usuarios     []Usuario `json:"usuarios,omitempty" gorm:"foreignKey:EmpresaID"`
	CriadoEm     time.Time `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time `json:"atualizadoEm" gorm:"not null"`
}
