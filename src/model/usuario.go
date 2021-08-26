package model

import (
	"cookbook/src/constants"
	"time"
)

// Usuario ... usuario Database Model
type Usuario struct {
	ID           uint64                  `json:"usuarioId,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Email        string                  `json:"email,omitempty" binding:"required,email" gorm:"type:varchar(100);not null;unique"`
	Senha        string                  `json:"senha,omitempty" binding:"required" gorm:"not null"`
	Nome         string                  `json:"nome,omitempty" binding:"required" gorm:"not null"`
	EmpresaID    uint64                  `json:"empresaId,omitempty" gorm:"not null"`
	Perfil       constants.PerfilUsuario `json:"perfil,omitempty" binding:"required" gorm:"not null"`
	CriadoEm     time.Time               `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time               `json:"atualizadoEm" gorm:"not null"`
}
