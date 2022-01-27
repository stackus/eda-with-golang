package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/stores/storespb"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	storespb.RegisterStoresServiceServer(registrar, server{app: app})
	return nil
}
