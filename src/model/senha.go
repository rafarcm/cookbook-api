package model

type Senha struct {
	Atual string `json:"atual,omitempty" binding:"required"`
	Nova  string `json:"nova,omitempty" binding:"required"`
}
