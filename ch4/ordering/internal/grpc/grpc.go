package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/ordering/internal/application"
	"github.com/stackus/eda-with-golang/ch4/ordering/orderingpb"
)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	orderingpb.RegisterOrderingServiceServer(registrar, server{app: app})
	return nil
}
