package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/stackus/errors"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application/ports"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/domain"
)

type OfferingRepository struct {
	tableName string
	db        *sql.DB
}

var _ ports.OfferingRepository = (*OfferingRepository)(nil)

func NewOfferingRepository(tableName string, db *sql.DB) OfferingRepository {
	return OfferingRepository{tableName: tableName, db: db}
}

func (r OfferingRepository) FindOffering(ctx context.Context, id, storeID string) (*domain.Offering, error) {
	const query = "SELECT name, description, price FROM %s WHERE id = $1 AND store_id = $2 LIMIT 1"

	offering := &domain.Offering{
		ID:      id,
		StoreID: storeID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), id, storeID).Scan(&offering.Name, &offering.Description, &offering.Price)
	if err != nil {
		return nil, errors.Wrap(err, "scanning offering")
	}

	return offering, nil
}

func (r OfferingRepository) AddOffering(ctx context.Context, offering *domain.Offering) error {
	const query = "INSERT INTO %s (id, store_id, name, description, price) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.db.ExecContext(ctx, r.table(query), offering.ID, offering.StoreID, offering.Name, offering.Description, offering.Price)

	return errors.Wrap(err, "inserting offering")
}

func (r OfferingRepository) RemoveOffering(ctx context.Context, id, storeID string) error {
	const query = "DELETE FROM %s WHERE id = $1 AND store_id = $2 LIMIT 1"

	_, err := r.db.ExecContext(ctx, r.table(query), id, storeID)

	return errors.Wrap(err, "deleting offering")
}

func (r OfferingRepository) GetStoreOfferings(ctx context.Context, storeID string) ([]*domain.Offering, error) {
	const query = "SELECT id, name, description, price FROM %s WHERE store_id = $1"

	offerings := make([]*domain.Offering, 0)

	rows, err := r.db.QueryContext(ctx, r.table(query), storeID)
	if err != nil {
		return nil, errors.Wrap(err, "querying store offerings")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing store offering rows")
		}
	}(rows)

	for rows.Next() {
		offering := &domain.Offering{
			StoreID: storeID,
		}
		err := rows.Scan(&offering.ID, &offering.Name, &offering.Description, &offering.Price)
		if err != nil {
			return nil, errors.Wrap(err, "scanning offering")
		}

		offerings = append(offerings, offering)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing offering rows")
	}

	return offerings, nil
}

func (r OfferingRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
