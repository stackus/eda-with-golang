package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/stackus/errors"

	"eda-in-golang/ch7/stores/internal/domain"
)

type CatalogRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.CatalogRepository = (*CatalogRepository)(nil)

func NewCatalogRepository(tableName string, db *sql.DB) CatalogRepository {
	return CatalogRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r CatalogRepository) AddProduct(ctx context.Context, productID, storeID, name, description, sku string,
	price float64,
) error {
	const query = `INSERT INTO %s (id, store_id, name, description, sku, price) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, r.table(query), productID, storeID, name, description, sku, price)

	return err
}

func (r CatalogRepository) Rebrand(ctx context.Context, productID, name, description string) error {
	const query = `UPDATE %s SET name = $2, description = $3 WHERE id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), productID, name, description)

	return err
}

func (r CatalogRepository) UpdatePrice(ctx context.Context, productID string, price float64) error {
	const query = `UPDATE %s SET price = $2 WHERE id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), productID, price)

	return err
}

func (r CatalogRepository) RemoveProduct(ctx context.Context, productID string) error {
	const query = `DELETE FROM %s WHERE id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), productID)

	return err
}

func (r CatalogRepository) Find(ctx context.Context, productID string) (*domain.Product, error) {
	const query = `SELECT store_id, name, description, sku, price FROM %s WHERE id = $1 LIMIT 1`

	product := domain.NewProduct(productID)

	err := r.db.QueryRowContext(ctx, r.table(query), productID).Scan(&product.StoreID, &product.Name, &product.Description, &product.SKU, &product.Price)
	if err != nil {
		return nil, errors.Wrap(err, "scanning product")
	}

	return product, nil
}

func (r CatalogRepository) GetCatalog(ctx context.Context, storeID string) ([]*domain.Product, error) {
	const query = `SELECT id, name, description, sku, price FROM %s WHERE store_id = $1`

	products := make([]*domain.Product, 0)

	rows, err := r.db.QueryContext(ctx, r.table(query), storeID)
	if err != nil {
		return nil, errors.Wrap(err, "querying products")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing product rows")
		}
	}(rows)

	for rows.Next() {
		var productID, name, description, sku string
		var price float64
		err := rows.Scan(&productID, &name, &description, &sku, &price)
		if err != nil {
			return nil, errors.Wrap(err, "scanning product")
		}

		product := domain.NewProduct(productID)
		product.StoreID = storeID
		product.Name = name
		product.Description = description
		product.SKU = sku
		product.Price = price

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing product rows")
	}

	return products, nil
}

func (r CatalogRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
