package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type StoreRepository struct {
	tableName string
	db        *sql.DB
}

func NewStoreRepository(tableName string, db *sql.DB) StoreRepository {
	return StoreRepository{tableName: tableName, db: db}
}

var _ ports.StoreRepository = (*StoreRepository)(nil)

func (r StoreRepository) FindStore(ctx context.Context, storeID string) (*domain.Store, error) {
	const query = "SELECT name, location, participating FROM %s WHERE id = $1 LIMIT 1"

	store := new(domain.Store)

	var location string

	row := r.db.QueryRowContext(ctx, r.fmtQuery(query), storeID)
	err := row.Scan(&store.Name, location, &store.Participating)
	if err != nil {
		return nil, err
	}

	store.Location = domain.NewLocation(location)

	return store, nil
}

func (r StoreRepository) SaveStore(ctx context.Context, store *domain.Store) error {
	const query = "INSERT INTO %s (id, name, location, participating) VALUES ($1, $2, $3, $4)"

	_, err := r.db.ExecContext(ctx, r.fmtQuery(query), store.ID, store.Name, store.Location.String())

	return err
}

func (r StoreRepository) UpdateStore(ctx context.Context, store *domain.Store) error {
	const query = "UPDATE %s SET name = $1, location = $2, participating = $3 WHERE id = $4"

	_, err := r.db.ExecContext(ctx, r.fmtQuery(query), store.Name, store.Location.String(), store.Participating, store.ID)

	return err
}

func (r StoreRepository) fmtQuery(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
