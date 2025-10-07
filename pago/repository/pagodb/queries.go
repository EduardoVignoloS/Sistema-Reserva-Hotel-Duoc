package pagodb

var (
	createPayment = `
	INSERT INTO pago (id_reserva, monto, metodo, fecha_pago)
	VALUES (:id_reserva, :monto, :metodo, :fecha_pago)
	`
)
