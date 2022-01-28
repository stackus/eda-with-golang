package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/stackus/eda-with-golang/ch4/stores/storespb"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
)

func RegisterGateway(ctx context.Context, _ application.App, mux *chi.Mux, grpcAddr string) error {
	const apiRoot = "/api/stores"
	const specRoot = "/stores-spec/"

	gateway := runtime.NewServeMux()
	err := storespb.RegisterStoresServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// mount the GRPC gateway
	mux.Mount(apiRoot, gateway)
	// mount the swagger specifications
	mux.Mount(specRoot, http.StripPrefix(specRoot, http.FileServer(http.FS(swaggerUI))))

	return nil
}
