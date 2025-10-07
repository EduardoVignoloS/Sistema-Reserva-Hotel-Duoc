package reserva

import (
	"context"
	"fmt"
	"time"
)

// type service is the implementation of Service interface containing all the business logic
// and dependencies required to complete the given tasks without exposing the implementation.
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) CreateReservation(ctx context.Context, reserva Reserva) error {
	now := time.Now().Format("2006-01-02")
	fmt.Println(reserva.IDCliente, "id clinete")
	fmt.Printf("%+v", reserva)
	reserva.FechaReserva = now
	if err := s.repository.CreateReservation(ctx, reserva); err != nil {
		return fmt.Errorf("error creating reservation: %w", err)
	}

	return nil
}

func (s *service) GetReservationByID(ctx context.Context, id int) ([]Reserva, error) {
	//Obtenemos las reservas por usuario
	reservas, err := s.repository.GetReservationByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting reservation by ID: %w", err)
	}

	return reservas, nil
}

func (s *service) UpdateReservation(ctx context.Context, reserva Reserva) error {
	if err := s.repository.UpdateReservation(ctx, reserva); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteReservation(ctx context.Context, id int) error {
	if err := s.repository.DeleteReservation(ctx, id); err != nil {
		return err
	}
	return nil
}
func (s *service) ListReservations(ctx context.Context) ([]Reserva, error) {
	reservas, err := s.repository.ListReservations(ctx)
	if err != nil {
		return nil, err
	}
	return reservas, nil
}
