package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/proto/storespb"
)

func RegisterGateway(ctx context.Context, _ application.App, mux *chi.Mux, grpcAddr string) error {
	const storeAPIRoot = "/api/stores"
	const storeUIRoot = "/stores-ui/"

	gateway := runtime.NewServeMux()
	err := storespb.RegisterStoresServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// mount the stores GRPC gateway
	mux.Mount(storeAPIRoot, gateway)
	// mount the stores GRPC-Gateway swagger UI
	mux.Mount(storeUIRoot, http.StripPrefix(storeUIRoot, http.FileServer(http.FS(swaggerUI))))

	return nil
}
