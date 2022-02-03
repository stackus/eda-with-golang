package postgres

import (
	"context"
	"database/sql"
)

type DB interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type unitOfWork struct {
	db *sql.DB
}

// This should return a shared interface
func (w unitOfWork) getDb(ctx context.Context) DB {
	if tx := ctx.Value(dbContextKey); tx != nil {
		return tx.(*sql.Tx)
	}

	return w.db
}

const dbContextKey string = "postgres"

func initPostgres(db *sql.DB) UnitOfWorkMiddleware {
	return func(next UnitOfWorkFunc) UnitOfWorkFunc {
		return func(ctx context.Context) (err error) {
			var tx *sql.Tx

			tx, err = db.Begin()
			if err != nil {
				return err
			}

			dbCtx := context.WithValue(ctx, dbContextKey, tx)

			// recover panics etc...

			err = next(dbCtx)

			return err
		}
	}

	// return func(ctx context.Context, next func(context.Context) error) (err error) {
	// 	tx, err := db.Begin()
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// 	dbCtx := context.WithValue(ctx, dbContextKey, tx)
	//
	// 	err = next(dbCtx)
	//
	// 	return err
	// }
}

// in application

type UnitOfWorkFunc func(context.Context) error
type UnitOfWorkMiddleware func(next UnitOfWorkFunc) UnitOfWorkFunc

type UnitOfWork struct {
	fns []UnitOfWorkMiddleware
}

func NewUnitOfWork(fns ...UnitOfWorkMiddleware) UnitOfWork {
	return UnitOfWork{fns: fns}
}

func (w UnitOfWork) Session(ctx context.Context, unit func(context.Context) error) error {
	h := unit
	for _, fn := range w.fns {
		h = fn(h)
	}

	return h(ctx)
}

func foo() {

}
