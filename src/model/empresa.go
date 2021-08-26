package model

import "time"

// Empresa ... empresa Database Model
type Empresa struct {
	ID           uint64        `json:"empresaId,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Nome         string        `json:"nome,omitempty" binding:"required" gorm:"not null"`
	Usuarios     []Usuario     `json:"usuarios,omitempty" binding:"dive" gorm:"foreignKey:EmpresaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Receitas     []Receita     `json:"receitas,omitempty" gorm:"foreignKey:EmpresaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ingredientes []Ingrediente `json:"ingredientes,omitempty" gorm:"foreignKey:EmpresaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Utensilios   []Utensilio   `json:"utensilios,omitempty" gorm:"foreignKey:EmpresaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CriadoEm     time.Time     `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time     `json:"atualizadoEm" gorm:"not null"`
}
