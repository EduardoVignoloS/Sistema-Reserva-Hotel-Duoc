package pagodb

type PaymentDB struct {
	IdPago    int     `db:"id_pago"`
	IdReserva int     `db:"id_reserva"`
	Monto     float64 `db:"monto"`
	FechaPago string  `db:"fecha_pago"`
	TipoPago  string  `db:"tipo_pago"`
	Estado    string  `db:"estado"`
}
