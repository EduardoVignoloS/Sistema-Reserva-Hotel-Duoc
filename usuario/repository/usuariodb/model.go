package usuariodb

import "github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/usuario"

type UsuarioDB struct {
	ID             string `db:"id_cliente"`
	Nombre         string `db:"nombre"`
	Apellido       string `db:"apellido"`
	Email          string `db:"email"`
	Password       string `db:"password"`
	Telefono       string `db:"telefono"`
	TypeC          string `db:"typeC"`
	Fecha_Registro string `db:"fecha_registro"`
}

func toUsuarioDB(u UsuarioDB) usuario.Usuario {
	return usuario.Usuario{
		ID:             u.ID,
		Nombre:         u.Nombre,
		Apellido:       u.Apellido,
		Email:          u.Email,
		Password:       u.Password,
		Telefono:       u.Telefono,
		TypeC:          u.TypeC,
		Fecha_Registro: u.Fecha_Registro,
	}
}
