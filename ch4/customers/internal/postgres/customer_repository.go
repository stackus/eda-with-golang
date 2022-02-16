package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/stackus/eda-with-golang/ch4/customers/internal/domain"
)

type CustomerRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.CustomerRepository = (*CustomerRepository)(nil)

func NewCustomerRepository(tableName string, db *sql.DB) CustomerRepository {
	return CustomerRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	const query = ""

	// TODO implement me
	panic("implement me")
}

func (r CustomerRepository) Find(ctx context.Context, customerID domain.CustomerID) (*domain.Customer, error) {
	const query = ""

	// TODO implement me
	panic("implement me")
}

func (r CustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	const query = ""

	// TODO implement me
	panic("implement me")
}

func (r CustomerRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
