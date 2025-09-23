package usuariodb

import (
	"context"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/usuario"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateAccount(ctx context.Context, usuario usuario.Usuario) error {
	data := map[string]any{
		"id_cliente":     usuario.ID,
		"nombre":         usuario.Nombre,
		"apellido":       usuario.Apellido,
		"email":          usuario.Email,
		"password":       usuario.Password,
		"telefono":       usuario.Telefono,
		"typeC":          usuario.TypeC,
		"fecha_registro": usuario.Fecha_Registro,
	}
	_, err := r.db.NamedExecContext(ctx, `INSERT INTO cliente (id_cliente, nombre, apellido, email, password, telefono, typeC, fecha_registro) 
	VALUES (:id_cliente, :nombre, :apellido, :email, :password, :telefono, :typeC, :fecha_registro)`, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Login(ctx context.Context, usuarioCore usuario.Usuario) (usuario.Usuario, error) {
	var usuarioDB UsuarioDB

	query := `SELECT * FROM cliente WHERE email = :email AND password = :password`
	if err := r.db.GetContext(ctx, &usuarioDB, query, usuarioCore.Email, usuarioCore.Password); err != nil {
		return usuario.Usuario{}, err
	}

	return toUsuarioDB(usuarioDB), nil
}
