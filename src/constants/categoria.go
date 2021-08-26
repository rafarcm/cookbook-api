package constants

type Categoria uint8

const (
	Acompanhamentos Categoria = iota + 1
	BalasDocinhos
	Bebidas
	Biscoitos
	BolosRecheios
	Carnes
	DocesSobremessas
	FrutosDoMar
	Geleias
	Lanches
	Massas
	Molhos
	PaesSalgados
	PaesDoces
	Petiscos
	Pizzas
	Pratos
	Risotos
	Saladas
	Salgadinhos
	Sopas
)

func (c Categoria) String() string {
	return [...]string{
		"Acompanhamentos",
		"Balas e Docinhos",
		"Bebidas",
		"Biscoitos",
		"Bolos e Recheios",
		"Carnes",
		"Doces e Sobremessas",
		"Frutos do Mar",
		"Geleias",
		"Lanches",
		"Massas",
		"Molhos",
		"Pães Salgados",
		"Pães Doces",
		"Petiscos",
		"Pizzas",
		"Pratos",
		"Risotos",
		"Saladas",
		"Salgadinhos",
		"Sopas",
	}[c]
}
