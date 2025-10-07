package pago

import (
	"context"
)

type Service interface {
	CreatePayment(ctx context.Context, payment Payment) error
}

type Repository interface {
	CreatePayment(ctx context.Context, payment Payment) error
}
