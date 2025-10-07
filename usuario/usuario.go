package usuario

import "context"

// Service interface defines the set of functions that will containt the business logic, allowing
// CRUD operations and more.
type Service interface {
	CreateAccount(ctx context.Context, usuario Usuario) error
	Login(ctx context.Context, usuario Usuario) (Usuario, error)
}

type Repository interface {
	CreateAccount(ctx context.Context, usuario Usuario) error
	Login(ctx context.Context, usuario Usuario) (Usuario, error)
	Query(ctx context.Context, email string) (string, error)
}
