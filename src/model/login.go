package model

// Login ... Login Database Model
type Login struct {
	Email string `json:"email,omitempty" binding:"required"`
	Senha string `json:"senha,omitempty" binding:"required"`
}
