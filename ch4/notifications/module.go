package notifications

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
	"github.com/stackus/eda-with-golang/ch4/notifications/internal/application"
	"github.com/stackus/eda-with-golang/ch4/notifications/internal/grpc"
	"github.com/stackus/eda-with-golang/ch4/notifications/internal/logging"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)

	var app application.App
	app = application.New(customers)
	app = logging.NewApplication(app, mono.Logger())

	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
