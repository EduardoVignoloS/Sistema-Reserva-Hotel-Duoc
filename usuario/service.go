package usuario

import "context"

// type service is the implementation of Service interface containing all the business logic
// and dependencies required to complete the given tasks without exposing the implementation.
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) CreateAccount(ctx context.Context, usuario Usuario) error {
	if err := s.repository.CreateAccount(ctx, usuario); err != nil {
		return err
	}

	return nil
}
func (s *service) Login(ctx context.Context, usuario Usuario) (Usuario, error) {
	usuario, err := s.repository.Login(ctx, usuario)
	if err != nil {
		return Usuario{}, err
	}

	return usuario, nil
}
