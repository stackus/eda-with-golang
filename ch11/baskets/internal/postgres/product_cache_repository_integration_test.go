//go:build integration

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"eda-in-golang/baskets/internal/domain"
	"eda-in-golang/internal/logger/log"
	"eda-in-golang/migrations"
)

type productCacheSuite struct {
	container testcontainers.Container
	db        *sql.DB
	mock      *domain.MockProductRepository
	repo      ProductCacheRepository
	suite.Suite
}

func TestProductCacheRepository(t *testing.T) {
	suite.Run(t, &productCacheSuite{})
}

func (s *productCacheSuite) SetupSuite() {
	var err error

	ctx := context.Background()
	initDir, err := filepath.Abs("./../../../docker/database")
	if err != nil {
		s.T().Fatal(err)
	}
	s.container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:12-alpine",
			Hostname:     "postgres",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_PASSWORD": "itsasecret",
			},
			Mounts: []testcontainers.ContainerMount{
				testcontainers.BindMount(initDir, "/docker-entrypoint-initdb.d"),
			},
			WaitingFor: wait.ForSQL("5432/tcp", "pgx", func(port nat.Port) string {
				return fmt.Sprintf("postgres://mallbots_user:mallbots_pass@localhost:%s/mallbots?sslmode=disable", port.Port())
			}).Timeout(5 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		s.T().Fatal(err)
	}

	endpoint, err := s.container.Endpoint(ctx, "")
	if err != nil {
		s.T().Fatal(err)
	}

	s.db, err = sql.Open("pgx", fmt.Sprintf("postgres://mallbots_user:mallbots_pass@%s/mallbots?sslmode=disable", endpoint))
	if err != nil {
		s.T().Fatal(err)
	}

	goose.SetLogger(&log.SilentLogger{})
	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		s.T().Fatal(err)
	}
	if err := goose.Up(s.db, "."); err != nil {
		s.T().Fatal(err)
	}
}
func (s *productCacheSuite) TearDownSuite() {
	err := s.db.Close()
	if err != nil {
		s.T().Fatal(err)
	}
	if err := s.container.Terminate(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func (s *productCacheSuite) SetupTest() {
	_, err := s.db.ExecContext(context.Background(), "TRUNCATE baskets.products_cache")
	if err != nil {
		s.T().Fatal(err)
	}
	s.mock = domain.NewMockProductRepository(s.T())
	s.repo = NewProductCacheRepository("baskets.products_cache", s.db, s.mock)
}
func (s *productCacheSuite) TearDownTest() {
}

func (s *productCacheSuite) TestProductCacheRepository_Add() {
	s.Assert().NoError(s.repo.Add(context.Background(), "product-id", "store-id", "product-name", 10.00))
	row := s.db.QueryRow("SELECT name FROM baskets.products_cache WHERE id = $1", "product-id")
	if s.Assert().NoError(row.Err()) {
		var name string
		s.Assert().NoError(row.Scan(&name))
		s.Assert().Equal("product-name", name)
	}
}

func (s *productCacheSuite) TestProductCacheRepository_AddDupe() {
	s.Assert().NoError(s.repo.Add(context.Background(), "product-id", "store-id", "product-name", 10.00))
	s.Assert().NoError(s.repo.Add(context.Background(), "product-id", "store-id", "dupe-product-name", 10.00))
	row := s.db.QueryRow("SELECT name FROM baskets.products_cache WHERE id = $1", "product-id")
	if s.Assert().NoError(row.Err()) {
		var name string
		s.Assert().NoError(row.Scan(&name))
		s.Assert().Equal("product-name", name)
	}
}

func (s *productCacheSuite) TestProductCacheRepository_Rebrand() {
	_ = s.repo.Add(context.Background(), "product-id", "store-id", "product-name", 10.00)

	s.Assert().NoError(s.repo.Rebrand(context.Background(), "product-id", "new-product-name"))
	row := s.db.QueryRow("SELECT name FROM baskets.products_cache WHERE id = $1", "product-id")
	if s.Assert().NoError(row.Err()) {
		var name string
		s.Assert().NoError(row.Scan(&name))
		s.Assert().Equal("new-product-name", name)
	}
}

func (s *productCacheSuite) TestProductCacheRepository_UpdatePrice() {
	_ = s.repo.Add(context.Background(), "product-id", "store-id", "product-name", 10.00)

	s.Assert().NoError(s.repo.UpdatePrice(context.Background(), "product-id", 2.00))
	row := s.db.QueryRow("SELECT price FROM baskets.products_cache WHERE id = $1", "product-id")
	if s.Assert().NoError(row.Err()) {
		var price float64
		s.Assert().NoError(row.Scan(&price))
		s.Assert().Equal(12.00, price)
	}
}
func (s *productCacheSuite) TestProductCacheRepository_Remove() {
	_ = s.repo.Add(context.Background(), "product-id", "store-id", "product-name", 10.00)

	s.Assert().NoError(s.repo.Remove(context.Background(), "product-id"))
	row := s.db.QueryRow("SELECT name FROM baskets.products_cache WHERE id = $1", "product-id")
	if s.Assert().NoError(row.Err()) {
		var name string
		s.Assert().Error(row.Scan(&name))
	}
}

func (s *productCacheSuite) TestProductCacheRepository_Find() {
	_ = s.repo.Add(context.Background(), "product-id", "store-id", "product-name", 10.00)

	product, err := s.repo.Find(context.Background(), "product-id")
	if s.Assert().NoError(err) {
		s.Assert().NotNil(product)
		s.Assert().Equal("product-name", product.Name)
	}
}

func (s *productCacheSuite) TestProductCacheRepository_FindFromFallback() {
	s.mock.On("Find", mock.Anything, "product-id").Return(&domain.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}, nil)

	product, err := s.repo.Find(context.Background(), "product-id")
	if s.Assert().NoError(err) {
		s.Assert().NotNil(product)
		s.Assert().Equal("product-name", product.Name)
	}
}
