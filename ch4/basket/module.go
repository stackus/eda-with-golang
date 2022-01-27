package basket

import (
	"context"

	rpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/stackus/eda-with-golang/ch4/basket/internal/application"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/grpc"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/postgres"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/rest"
	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// Startup Driven adapters
	basketRepo := postgres.NewBasketRepository("basket.baskets", mono.DB())
	conn, err := rpc.Dial(mono.Config().Rpc.Address(), rpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	productRepo := grpc.NewProductRepository(conn)
	defer func() {
		if err != nil {
			if err = conn.Close(); err != nil {
				// TODO do something when logging is a thing
			}
			return
		}
		go func() {
			<-ctx.Done()
			if err = conn.Close(); err != nil {
				// TODO do something when logging is a thing
			}
		}()
	}()

	// Startup application
	app := application.New(basketRepo, productRepo)

	// Setup Driver adapters
	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, app, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}

	return nil
}
