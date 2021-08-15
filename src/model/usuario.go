package model

import (
	"time"
)

// Usuario ... usuario Database Model
type Usuario struct {
	ID           uint64    `json:"usuarioId,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Username     string    `json:"username,omitempty" binding:"required" gorm:"type:varchar(100);not null;unique"`
	Senha        string    `json:"senha,omitempty" binding:"required" gorm:"not null"`
	EmpresaID    uint64    `json:"empresaId,omitempty" binding:"required" gorm:"not null"`
	CriadoEm     time.Time `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time `json:"atualizadoEm" gorm:"not null"`
}
