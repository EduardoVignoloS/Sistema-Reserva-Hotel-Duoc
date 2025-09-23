package reserva

// Service interface defines the set of functions that will containt the business logic, allowing
// CRUD operations and more.
type Service interface {
	CreateReservation(reserva Reserva) error
	GetReservationByID(id int) (Reserva, error)
	UpdateReservation(reserva Reserva) error
	DeleteReservation(id int) error
	ListReservations() ([]Reserva, error)
}

type Repository interface {
	CreateReservation(reserva Reserva) error
	GetReservationByID(id int) (Reserva, error)
	UpdateReservation(reserva Reserva) error
	DeleteReservation(id int) error
	ListReservations() ([]Reserva, error)
}
