package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/baskets/basketspb"
	"github.com/stackus/eda-with-golang/ch4/baskets/internal/application"
)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	basketspb.RegisterBasketServiceServer(registrar, server{app: app})
	return nil
}
