package postgres

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"

	"eda-in-golang/customers/internal/domain"
	"eda-in-golang/internal/postgres"
)

type CustomerRepository struct {
	tableName string
	db        postgres.DB
}

var _ domain.CustomerRepository = (*CustomerRepository)(nil)

func NewCustomerRepository(tableName string, db postgres.DB) CustomerRepository {
	return CustomerRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r CustomerRepository) Find(ctx context.Context, customerID string) (*domain.Customer, error) {
	const query = "SELECT name, sms_number, enabled FROM %s WHERE id = $1 LIMIT 1"

	ctx, span := tracer.Start(ctx, "Find")

	tableQuery := r.table(query)

	span.SetAttributes(
		attribute.String("Query", tableQuery),
	)

	customer := domain.NewCustomer(customerID)

	err := r.db.QueryRowContext(ctx, tableQuery, customerID).Scan(&customer.Name, &customer.SmsNumber, &customer.Enabled)

	return customer, err
}

func (r CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	const query = "INSERT INTO %s (id, name, sms_number, enabled) VALUES ($1, $2, $3, $4)"

	ctx, span := tracer.Start(ctx, "Save")

	tableQuery := r.table(query)

	span.SetAttributes(
		attribute.String("Exec", tableQuery),
	)

	_, err := r.db.ExecContext(ctx, tableQuery, customer.ID(), customer.Name, customer.SmsNumber, customer.Enabled)

	return err
}

func (r CustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	const query = "UPDATE %s SET name = $2, sms_number = $3, enabled = $4 WHERE id = $1"

	ctx, span := tracer.Start(ctx, "Update")

	tableQuery := r.table(query)

	span.SetAttributes(
		attribute.String("Exec", tableQuery),
	)

	_, err := r.db.ExecContext(ctx, tableQuery, customer.ID(), customer.Name, customer.SmsNumber, customer.Enabled)

	return err
}

func (r CustomerRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
