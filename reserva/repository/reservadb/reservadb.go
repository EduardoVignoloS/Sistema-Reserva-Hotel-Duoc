package reservadb

import (
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

func (r *Repository) CreateReservation(reserva reserva.Reserva) error {
	data := map[string]any{
		"id_cliente":    reserva.IDCliente,
		"id_habitacion": reserva.IDHabitacion,
		"fecha_inicio":  reserva.FechaInicio,
		"fecha_fin":     reserva.FechaFin,
		"total":         reserva.Total,
		"estado":        reserva.Estado,
		"fecha_reserva": reserva.FechaReserva,
	}
	_, err := r.db.NamedExec(`INSERT INTO reserva (id_cliente, id_habitacion, fecha_inicio, fecha_fin, total, estado, fecha_reserva) 
	VALUES (:id_cliente, :id_habitacion, :fecha_inicio, :fecha_fin, :total, :estado, :fecha_reserva)`, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetReservationByID(id int) (reserva.Reserva, error) {
	var reservaDB ReservaDB

	query := `SELECT * FROM reserva WHERE id_reserva = $1`
	if err := r.db.Get(&reservaDB, query, id); err != nil {
		return reserva.Reserva{}, err
	}

	return reservaDB.ToReserva(), nil
}

func (r *Repository) UpdateReservation(reserva reserva.Reserva) error {
	data := map[string]any{
		"id_reserva":    reserva.IDReserva,
		"id_cliente":    reserva.IDCliente,
		"id_habitacion": reserva.IDHabitacion,
		"fecha_inicio":  reserva.FechaInicio,
		"fecha_fin":     reserva.FechaFin,
		"total":         reserva.Total,
		"estado":        reserva.Estado,
		"fecha_reserva": reserva.FechaReserva,
	}
	_, err := r.db.NamedExec(`UPDATE reserva SET id_cliente=:id_cliente, id_habitacion=:id_habitacion, fecha_inicio=:fecha_inicio, 
	fecha_fin=:fecha_fin, total=:total, estado=:estado, fecha_reserva=:fecha_reserva WHERE id_reserva=:id_reserva`, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteReservation(id int) error {
	_, err := r.db.Exec(`DELETE FROM reserva WHERE id_reserva = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListReservations() ([]reserva.Reserva, error) {
	var reservasDB []ReservaDB

	query := `SELECT * FROM reserva`
	if err := r.db.Select(&reservasDB, query); err != nil {
		return nil, err
	}

	reservas := make([]reserva.Reserva, len(reservasDB))
	for i, rDB := range reservasDB {
		reservas[i] = rDB.ToReserva()
	}

	return reservas, nil
}
