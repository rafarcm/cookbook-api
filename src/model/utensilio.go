package model

import (
	"fmt"
	"net/http"
	"time"
)

const (
	UtensilioGetAllRel  = "utensilio_busca_todos"
	UtensilioGetByIdRel = "utensilio_busca_id"
	UtensilioAddRel     = "utensilio_inclusao"
	UtensilioUpdateRel  = "utensilio_atualizacao"
	UtensilioDeleteRel  = "utensilio_exclusao"
)

// Utensilio ... utensilio Database Model
type Utensilio struct {
	ID           uint64    `json:"utensilioId,omitempty" gorm:"primaryKey;autoIncrement;unique;not null"`
	Descricao    string    `json:"descricao,omitempty" binding:"required" gorm:"not null"`
	UsuarioID    uint64    `json:"usuarioId,omitempty" gorm:"not null"`
	CriadoEm     time.Time `json:"criadoEm" gorm:"not null"`
	AtualizadoEm time.Time `json:"atualizadoEm" gorm:"not null"`
	Links        []Link    `json:"_links" gorm:"-"`
}

func (u *Utensilio) AddLinks(host string, rel string) {
	links := []Link{
		{
			Type:      http.MethodGet,
			Rel:       getRel(UtensilioGetAllRel, rel),
			URI:       host + "/utensilios{?descricao,usuarioId}",
			Descricao: "Retorna todos os utensílios",
		},
		{
			Type:      http.MethodGet,
			Rel:       getRel(UtensilioGetByIdRel, rel),
			URI:       host + fmt.Sprintf("/utensilios/%d", u.ID),
			Descricao: "Retorna um utensílio específico",
		},
		{
			Type:      http.MethodPost,
			Rel:       getRel(UtensilioAddRel, rel),
			URI:       host + "/utensilios",
			Descricao: "Adiciona um novo utensílio",
		},
		{
			Type:      http.MethodPut,
			Rel:       getRel(UtensilioUpdateRel, rel),
			URI:       host + fmt.Sprintf("/utensilios/%d", u.ID),
			Descricao: "Altera um utensílio específico",
		},
		{
			Type:      http.MethodDelete,
			Rel:       getRel(UtensilioDeleteRel, rel),
			URI:       host + fmt.Sprintf("/utensilios/%d", u.ID),
			Descricao: "Exclui um utensílio específico",
		},
	}

	u.Links = links
}

func getRel(rel1 string, rel2 string) string {
	if rel1 == rel2 {
		return "self"
	}

	return rel1
}
