package notifications

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
	"github.com/stackus/eda-with-golang/ch4/notifications/internal/application"
	"github.com/stackus/eda-with-golang/ch4/notifications/internal/grpc"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	app := application.New()

	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
