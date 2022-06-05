package grpc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/stackus/errors"
	"google.golang.org/grpc"

	"eda-in-golang/customers/customerspb"
	"eda-in-golang/customers/internal/application"
	"eda-in-golang/customers/internal/domain"
	"eda-in-golang/internal/di"
)

type server struct {
	c di.Container
	// app application.App
	customerspb.UnimplementedCustomersServiceServer
}

var _ customerspb.CustomersServiceServer = (*server)(nil)

func RegisterServer(container di.Container, registrar grpc.ServiceRegistrar) error {
	customerspb.RegisterCustomersServiceServer(registrar, server{c: container})
	return nil
}

func (s server) RegisterCustomer(ctx context.Context, request *customerspb.RegisterCustomerRequest) (resp *customerspb.RegisterCustomerResponse, err error) {
	ctx, cleanup := s.c.Scoped(ctx)
	defer cleanup()

	tx := di.Get(ctx, "tx").(*sql.Tx)
	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()
			panic(p)
		case err != nil:
			rErr := tx.Rollback()
			if rErr != nil {
				err = errors.Wrap(err, rErr.Error())
			}
		default:
			err = tx.Commit()
		}
	}()

	app := di.Get(ctx, "app").(application.App)

	id := uuid.New().String()
	err = app.RegisterCustomer(ctx, application.RegisterCustomer{
		ID:        id,
		Name:      request.GetName(),
		SmsNumber: request.GetSmsNumber(),
	})
	return &customerspb.RegisterCustomerResponse{Id: id}, err
}

func (s server) AuthorizeCustomer(ctx context.Context, request *customerspb.AuthorizeCustomerRequest) (resp *customerspb.AuthorizeCustomerResponse, err error) {
	ctx, cleanup := s.c.Scoped(ctx)
	defer cleanup()

	tx := di.Get(ctx, "tx").(*sql.Tx)
	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()
			panic(p)
		case err != nil:
			rErr := tx.Rollback()
			if rErr != nil {
				err = errors.Wrap(err, rErr.Error())
			}
		default:
			err = tx.Commit()
		}
	}()

	app := di.Get(ctx, "app").(application.App)

	err = app.AuthorizeCustomer(ctx, application.AuthorizeCustomer{
		ID: request.GetId(),
	})

	return &customerspb.AuthorizeCustomerResponse{}, err
}

func (s server) GetCustomer(ctx context.Context, request *customerspb.GetCustomerRequest) (resp *customerspb.GetCustomerResponse, err error) {
	ctx, cleanup := s.c.Scoped(ctx)
	defer cleanup()

	tx := di.Get(ctx, "tx").(*sql.Tx)
	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()
			panic(p)
		case err != nil:
			rErr := tx.Rollback()
			if rErr != nil {
				err = errors.Wrap(err, rErr.Error())
			}
		default:
			err = tx.Commit()
		}
	}()

	app := di.Get(ctx, "app").(application.App)

	customer, err := app.GetCustomer(ctx, application.GetCustomer{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &customerspb.GetCustomerResponse{
		Customer: s.customerFromDomain(customer),
	}, nil
}

func (s server) EnableCustomer(ctx context.Context, request *customerspb.EnableCustomerRequest) (resp *customerspb.EnableCustomerResponse, err error) {
	ctx, cleanup := s.c.Scoped(ctx)
	defer cleanup()

	tx := di.Get(ctx, "tx").(*sql.Tx)
	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()
			panic(p)
		case err != nil:
			rErr := tx.Rollback()
			if rErr != nil {
				err = errors.Wrap(err, rErr.Error())
			}
		default:
			err = tx.Commit()
		}
	}()

	app := di.Get(ctx, "app").(application.App)

	err = app.EnableCustomer(ctx, application.EnableCustomer{ID: request.GetId()})
	return &customerspb.EnableCustomerResponse{}, err
}

func (s server) DisableCustomer(ctx context.Context, request *customerspb.DisableCustomerRequest) (resp *customerspb.DisableCustomerResponse, err error) {
	ctx, cleanup := s.c.Scoped(ctx)
	defer cleanup()

	tx := di.Get(ctx, "tx").(*sql.Tx)
	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()
			panic(p)
		case err != nil:
			rErr := tx.Rollback()
			if rErr != nil {
				err = errors.Wrap(err, rErr.Error())
			}
		default:
			err = tx.Commit()
		}
	}()

	app := di.Get(ctx, "app").(application.App)

	err = app.DisableCustomer(ctx, application.DisableCustomer{ID: request.GetId()})
	return &customerspb.DisableCustomerResponse{}, err
}

func (s server) customerFromDomain(customer *domain.Customer) *customerspb.Customer {
	return &customerspb.Customer{
		Id:        customer.ID(),
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}
}
