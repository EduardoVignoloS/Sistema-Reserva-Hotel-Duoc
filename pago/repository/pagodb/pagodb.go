package pagodb

import (
	"context"
	"fmt"
	"time"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/kit/pgx"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/pago"
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

func (r *Repository) CreatePayment(ctx context.Context, payment pago.Payment) error {
	fmt.Println(payment, "payment en repo")
	data := map[string]any{
		"id_reserva": payment.IDReserva,
		"monto":      payment.Monto,
		"metodo":     payment.TipoPago,
		"fecha_pago": time.Now().Format("2006-01-02"),
	}

	query := pgx.ParseQuery(createPayment, data)
	if err := pgx.RunCUD(ctx, r.db, query, data); err != nil {
		return err
	}

	return nil
}
