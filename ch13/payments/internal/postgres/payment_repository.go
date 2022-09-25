package postgres

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"

	"eda-in-golang/internal/postgres"
	"eda-in-golang/payments/internal/application"
	"eda-in-golang/payments/internal/models"
)

type PaymentRepository struct {
	tableName string
	db        postgres.DB
}

var _ application.PaymentRepository = (*PaymentRepository)(nil)

func NewPaymentRepository(tableName string, db postgres.DB) PaymentRepository {
	return PaymentRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r PaymentRepository) Save(ctx context.Context, payment *models.Payment) error {
	const query = "INSERT INTO %s (id, customer_id, amount) VALUES ($1, $2, $3)"

	ctx, span := tracer.Start(ctx, "Save")

	tableQuery := r.table(query)

	span.SetAttributes(
		attribute.String("Exec", tableQuery),
	)

	_, err := r.db.ExecContext(ctx, tableQuery, payment.ID, payment.CustomerID, payment.Amount)

	return err
}

func (r PaymentRepository) Find(ctx context.Context, paymentID string) (*models.Payment, error) {
	const query = "SELECT customer_id, amount FROM %s WHERE id = $1 LIMIT 1"

	ctx, span := tracer.Start(ctx, "Find")

	tableQuery := r.table(query)

	span.SetAttributes(
		attribute.String("Query", tableQuery),
	)

	payment := &models.Payment{
		ID: paymentID,
	}

	err := r.db.QueryRowContext(ctx, tableQuery, paymentID).Scan(&payment.CustomerID, &payment.Amount)

	return payment, err
}

func (r PaymentRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
