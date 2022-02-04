package depot

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/depot/internal/application"
	"github.com/stackus/eda-with-golang/ch4/depot/internal/grpc"
	"github.com/stackus/eda-with-golang/ch4/depot/internal/postgres"
	"github.com/stackus/eda-with-golang/ch4/depot/internal/rest"
	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	listRepo := postgres.NewShoppingListRepository("depot.shopping_lists", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	storeRepo := grpc.NewStoreRepository(conn)
	productRepo := grpc.NewProductRepository(conn)

	app := application.New(listRepo, storeRepo, productRepo)

	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, app, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}

	return nil
}
