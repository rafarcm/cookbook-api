package model

import (
	"time"
)

type Usuario struct {
	ID           uint64    `json:"id,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Nick         string    `json:"nick,omitempty" binding:"required" gorm:"not null"`
	Senha        string    `json:"senha,omitempty" binding:"required" gorm:"not null"`
	EmpresaID    uint64    `json:"empresa,omitempty" binding:"required" gorm:"not null"`
	CriadoEm     time.Time `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time `json:"atualizadoEm" gorm:"not null"`
}
