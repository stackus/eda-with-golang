package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/depot/depotpb"
	"github.com/stackus/eda-with-golang/ch4/depot/internal/application"
)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	depotpb.RegisterDepotServiceServer(registrar, server{app: app})
	return nil
}
