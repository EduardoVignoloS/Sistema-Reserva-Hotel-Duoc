package usuariodb

import (
	"context"
	"database/sql"
	"time"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/kit/pgx"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/usuario"
	"github.com/jmoiron/sqlx"
)

const vacio = ""

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Query(ctx context.Context, email string) (string, error) {
	data := map[string]any{
		"email": email,
	}

	query := pgx.ParseQuery(query, data)

	var user UsuarioDB
	err := pgx.RunQuery(ctx, r.db, query, &user)
	if err != nil && err != sql.ErrNoRows {
		return vacio, err
	}
	return user.Email, nil

}

func (r *Repository) CreateAccount(ctx context.Context, usuario usuario.Usuario) error {

	data := map[string]any{
		"nombre":         usuario.Nombre,
		"apellido":       usuario.Apellido,
		"email":          usuario.Email,
		"password":       usuario.Password,
		"telefono":       usuario.Telefono,
		"typeC":          usuario.TypeC,
		"fecha_registro": time.Now().Format("2006-01-02"),
	}

	query := pgx.ParseQuery(createUser, data)
	if err := pgx.RunCUD(ctx, r.db, query, data); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Login(ctx context.Context, usuarioCore usuario.Usuario) (usuario.Usuario, error) {
	var usuarioDB UsuarioDB

	data := map[string]any{
		"email":    usuarioCore.Email,
		"password": usuarioCore.Password,
	}

	query := pgx.ParseQuery(login, data)
	if err := pgx.RunQuery(ctx, r.db, query, &usuarioDB); err != nil {
		return usuario.Usuario{}, err
	}

	return toUsuarioDB(usuarioDB), nil
}
