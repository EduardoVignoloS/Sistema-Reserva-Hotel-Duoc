package reservadb

var (
	createReservation = `
	INSERT INTO reserva (id_cliente, numero_habitacion, fecha_inicio, fecha_fin, fecha_reserva) 
	VALUES (:id_cliente, :numero_habitacion, :fecha_inicio, :fecha_fin, :fecha_reserva)
	`

	reservationQueryByUser = `
	SELECT * FROM reserva WHERE id_cliente = :id_cliente
	`

	updateReservation = `
	UPDATE reserva SET id_cliente=:id_cliente, id_habitacion=:id_habitacion, fecha_inicio=:fecha_inicio, 
	fecha_fin=:fecha_fin, total=:total, estado=:estado, fecha_reserva=:fecha_reserva WHERE id_reserva=:id_reserva
	`

	updateReservationStatus = `
	UPDATE reserva SET estado=:estado WHERE id_reserva=:id_reserva
	`

	deleteReservation = `
	DELETE FROM reserva WHERE id_reserva = $1
	`

	listReservations = `
	SELECT * FROM reserva
	`
)
