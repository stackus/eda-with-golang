package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/stackus/eda-with-golang/ch4/basket/basketpb"
	"github.com/stackus/eda-with-golang/ch4/basket/internal/application"
)

func RegisterGateway(ctx context.Context, _ application.App, mux *chi.Mux, grpcAddr string) error {
	const storeAPIRoot = "/api/basket"
	const storeUIRoot = "/basket-ui/"

	gateway := runtime.NewServeMux()
	err := basketpb.RegisterBasketServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// mount the basket GRPC gateway
	mux.Mount(storeAPIRoot, gateway)
	// mount the basket GRPC-Gateway swagger UI
	mux.Mount(storeUIRoot, http.StripPrefix(storeUIRoot, http.FileServer(http.FS(swaggerUI))))

	return nil
}
