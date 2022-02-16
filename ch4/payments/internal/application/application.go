package application

import (
	"context"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/payments/internal/models"
)

type (
	CreateInvoice struct {
		ID      string
		OrderID string
		Amount  float64
	}

	AdjustInvoice struct {
		ID     string
		Amount float64
	}

	PayInvoice struct {
		ID string
	}

	CancelInvoice struct {
		ID string
	}

	App interface {
		CreateInvoice(ctx context.Context, create CreateInvoice) error
		AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error
		PayInvoice(ctx context.Context, pay PayInvoice) error
		CancelInvoice(ctx context.Context, cancel CancelInvoice) error
	}

	Application struct {
		invoices models.InvoiceRepository
	}
)

var _ App = (*Application)(nil)

func New(invoices models.InvoiceRepository) *Application {
	return &Application{invoices: invoices}
}

func (a Application) CreateInvoice(ctx context.Context, create CreateInvoice) error {
	return a.invoices.Save(ctx, &models.Invoice{
		ID:      create.ID,
		OrderID: create.OrderID,
		Amount:  create.Amount,
		Status:  models.InvoicePending,
	})
}

func (a Application) AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error {
	invoice, err := a.invoices.Find(ctx, adjust.ID)
	if err != nil {
		return err
	}

	invoice.Amount = adjust.Amount

	return a.invoices.Update(ctx, invoice)
}

func (a Application) PayInvoice(ctx context.Context, pay PayInvoice) error {
	invoice, err := a.invoices.Find(ctx, pay.ID)
	if err != nil {
		return err
	}

	if invoice.Status != models.InvoicePending {
		return errors.Wrap(errors.ErrBadRequest, "invoice cannot be paid for")
	}

	invoice.Status = models.InvoicePaid

	return a.invoices.Update(ctx, invoice)
}

func (a Application) CancelInvoice(ctx context.Context, cancel CancelInvoice) error {
	invoice, err := a.invoices.Find(ctx, cancel.ID)
	if err != nil {
		return err
	}

	if invoice.Status != models.InvoicePending {
		return errors.Wrap(errors.ErrBadRequest, "invoice cannot be paid for")
	}

	invoice.Status = models.InvoiceCanceled

	return a.invoices.Update(ctx, invoice)
}
