package domain

import (
	"context"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer *Customer) error
	Find(ctx context.Context, customerID CustomerID) (*Customer, error)
	Update(ctx context.Context, customer *Customer) error
}
