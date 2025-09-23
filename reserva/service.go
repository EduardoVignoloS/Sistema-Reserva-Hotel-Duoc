package reserva

// type service is the implementation of Service interface containing all the business logic
// and dependencies required to complete the given tasks without exposing the implementation.
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) CreateReservation(reserva Reserva) error {
	if err := s.repository.CreateReservation(reserva); err != nil {
		return err
	}

	return nil
}

func (s *service) GetReservationByID(id int) (Reserva, error) {
	reserva, err := s.repository.GetReservationByID(id)
	if err != nil {
		return Reserva{}, err
	}

	return reserva, nil
}

func (s *service) UpdateReservation(reserva Reserva) error {
	if err := s.repository.UpdateReservation(reserva); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteReservation(id int) error {
	if err := s.repository.DeleteReservation(id); err != nil {
		return err
	}
	return nil
}
func (s *service) ListReservations() ([]Reserva, error) {
	reservas, err := s.repository.ListReservations()
	if err != nil {
		return nil, err
	}
	return reservas, nil
}
