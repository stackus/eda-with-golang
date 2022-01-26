package monolith

import (
	"context"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/internal/config"
	"github.com/stackus/eda-with-golang/ch4/internal/egress"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	Mux() *chi.Mux
	RPC() *grpc.Server
	Wait(...egress.WaitFunc) error
	Context() context.Context
	CancelFunc() context.CancelFunc
}

type Module interface {
	Startup(Monolith) error
}
