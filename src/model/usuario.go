package model

import (
	"time"
)

// Usuario ... usuario Database Model
type Usuario struct {
	ID           uint64    `json:"usuarioId,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Nick         string    `json:"nick,omitempty" binding:"required" gorm:"not null"`
	Senha        string    `json:"senha,omitempty" binding:"required" gorm:"not null"`
	EmpresaID    uint64    `json:"empresaId,omitempty" binding:"required" gorm:"not null"`
	CriadoEm     time.Time `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time `json:"atualizadoEm" gorm:"not null"`
}
