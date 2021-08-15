package model

// Login ... Login Database Model
type Login struct {
	Username string `json:"username,omitempty" binding:"required"`
	Senha    string `json:"senha,omitempty" binding:"required"`
}
