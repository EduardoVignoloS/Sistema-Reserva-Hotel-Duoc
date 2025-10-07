package pago

import (
	"context"
	"fmt"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/reserva"
)

type service struct {
	repository        Repository
	repositoryReserva reserva.Repository
}

func NewService(repo Repository, repositoryReserva reserva.Repository) Service {
	return &service{repository: repo, repositoryReserva: repositoryReserva}
}

func (s *service) CreatePayment(ctx context.Context, payment Payment) error {

	//Creamos el pago y en caso exitoso, modificamos el estado de la reserva a "Pagada"
	if err := s.repository.CreatePayment(ctx, payment); err != nil {
		return fmt.Errorf("error creating payment: %w", err)
	}

	if err := s.repositoryReserva.UpdateStatusReservation(ctx, reserva.Reserva{
		IDReserva: payment.IDReserva,
		Estado:    "confirmada",
	}); err != nil {
		return fmt.Errorf("error updating reservation status: %w", err)
	}

	return nil
}
