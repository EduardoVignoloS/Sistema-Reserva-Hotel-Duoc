package reserva

import "context"

// Service interface defines the set of functions that will containt the business logic, allowing
// CRUD operations and more.
type Service interface {
	CreateReservation(ctx context.Context, reserva Reserva) error
	GetReservationByID(ctx context.Context, id int) ([]Reserva, error)
	UpdateReservation(ctx context.Context, reserva Reserva) error
	DeleteReservation(ctx context.Context, id int) error
	ListReservations(ctx context.Context) ([]Reserva, error)
}

type Repository interface {
	CreateReservation(ctx context.Context, reserva Reserva) error
	GetReservationByID(ctx context.Context, id int) ([]Reserva, error)
	UpdateReservation(ctx context.Context, reserva Reserva) error
	DeleteReservation(ctx context.Context, id int) error
	ListReservations(ctx context.Context) ([]Reserva, error)
	UpdateStatusReservation(ctx context.Context, reserva Reserva) error
}
