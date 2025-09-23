package reserva

type Reserva struct {
	IDReserva    int     `json:"IDReserva"`
	IDCliente    int     `json:"IDCliente"`
	IDHabitacion int     `json:"IDHabitacion"`
	FechaInicio  string  `json:"FechaInicio"`
	FechaFin     string  `json:"FechaFin"`
	Total        float64 `json:"Total"`
	Estado       string  `json:"Estado"`
	FechaReserva string  `json:"FechaReserva"`
}
