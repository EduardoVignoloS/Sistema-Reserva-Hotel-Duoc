package reservadb

import (
	"context"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/kit/pgx"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/reserva"
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

func (r *Repository) CreateReservation(ctx context.Context, reserva reserva.Reserva) error {
	data := map[string]any{
		"id_cliente":        reserva.IDCliente,
		"numero_habitacion": reserva.IDHabitacion,
		"fecha_inicio":      reserva.FechaInicio,
		"fecha_fin":         reserva.FechaFin,
		"fecha_reserva":     reserva.FechaReserva,
	}
	query := pgx.ParseQuery(createReservation, data)
	if err := pgx.RunCUD(ctx, r.db, query, data); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetReservationByID(ctx context.Context, id int) ([]reserva.Reserva, error) {

	data := map[string]any{
		"id_cliente": id,
	}

	query := pgx.ParseQuery(reservationQueryByUser, data)

	var reservationsDB []ReservaDB
	if err := pgx.RunQuerySlice(ctx, r.db, query, &reservationsDB); err != nil {
		return nil, err
	}

	reservationsCore := toCore(reservationsDB)

	return reservationsCore, nil
}

func (r *Repository) UpdateStatusReservation(ctx context.Context, reserva reserva.Reserva) error {
	data := map[string]any{
		"id_reserva": reserva.IDReserva,
		"estado":     reserva.Estado,
	}
	query := pgx.ParseQuery(updateReservationStatus, data)
	if err := pgx.RunCUD(ctx, r.db, query, data); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateReservation(ctx context.Context, reserva reserva.Reserva) error {
	data := map[string]any{
		"id_reserva":        reserva.IDReserva,
		"id_cliente":        reserva.IDCliente,
		"numero_habitacion": reserva.IDHabitacion,
		"fecha_inicio":      reserva.FechaInicio,
		"fecha_fin":         reserva.FechaFin,
		"total":             reserva.Total,
		"estado":            reserva.Estado,
		"fecha_reserva":     reserva.FechaReserva,
	}
	_, err := r.db.NamedExec(updateReservation, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteReservation(ctx context.Context, id int) error {
	_, err := r.db.Exec(deleteReservation, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListReservations(ctx context.Context) ([]reserva.Reserva, error) {
	var reservasDB []ReservaDB

	if err := r.db.Select(&reservasDB, listReservations); err != nil {
		return nil, err
	}

	reservas := make([]reserva.Reserva, len(reservasDB))
	for i, rDB := range reservasDB {
		reservas[i] = rDB.ToReserva()
	}

	return reservas, nil
}
