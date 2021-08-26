package constants

type PerfilUsuario uint8

const (
	Administrador PerfilUsuario = iota + 1
	Usuario
)

func (p PerfilUsuario) String() string {
	return [...]string{
		"Administrador",
		"Usuario",
	}[p]
}
