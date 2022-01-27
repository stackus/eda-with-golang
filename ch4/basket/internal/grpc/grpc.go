package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/basket/basketpb"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/application"
)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	basketpb.RegisterBasketServiceServer(registrar, server{app: app})
	return nil
}
