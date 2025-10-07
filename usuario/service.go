package usuario

import (
	"context"
	"errors"
	"fmt"
)

// type service is the implementation of Service interface containing all the business logic
// and dependencies required to complete the given tasks without exposing the implementation.
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) CreateAccount(ctx context.Context, usuario Usuario) error {

	fmt.Println("usuario en service:", usuario)
	email, err := s.repository.Query(ctx, usuario.Email)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error checking existing user: %w", err)
	}

	if email != "" {
		return fmt.Errorf("Usuario con ese correo ya existe")
	}

	if err := s.repository.CreateAccount(ctx, usuario); err != nil {
		fmt.Println(err)
		return fmt.Errorf("error creating account: %w", err)
	}

	return nil
}
func (s *service) Login(ctx context.Context, usuario Usuario) (Usuario, error) {
	email, err := s.repository.Query(ctx, usuario.Email)
	if err != nil {
		return Usuario{}, fmt.Errorf("error checking existing user: %w", err)
	}

	if email == "" {
		return Usuario{}, fmt.Errorf("Usuario con ese correo no existe")
	}

	usuarioDB, err := s.repository.Login(ctx, usuario)
	if err != nil {
		return Usuario{}, err
	}

	if usuarioDB.Password != usuario.Password {
		return Usuario{}, errors.New("password incorrecto")
	}

	return usuarioDB, nil
}
