package reservadb

import "github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/reserva"

type ReservaDB struct {
	IDReserva    int     `db:"id_reserva"`
	IDCliente    int     `db:"id_cliente"`
	IDHabitacion int     `db:"numero_habitacion"`
	FechaInicio  string  `db:"fecha_inicio"`
	FechaFin     string  `db:"fecha_fin"`
	Total        float64 `db:"total"`
	Estado       string  `db:"estado"`
	FechaReserva string  `db:"fecha_reserva"`
}

func toCore(reservationsDB []ReservaDB) []reserva.Reserva {
	var reservationsCore []reserva.Reserva
	for _, r := range reservationsDB {
		reservationsCore = append(reservationsCore, r.ToReserva())
	}
	return reservationsCore
}

func (r *ReservaDB) ToReserva() reserva.Reserva {
	return reserva.Reserva{
		IDReserva:    r.IDReserva,
		IDCliente:    r.IDCliente,
		IDHabitacion: r.IDHabitacion,
		FechaInicio:  r.FechaInicio,
		FechaFin:     r.FechaFin,
		Total:        r.Total,
		Estado:       r.Estado,
		FechaReserva: r.FechaReserva,
	}
}
