package rest

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	grpcstd "google.golang.org/grpc"

	"eda-in-golang/baskets/internal/application"
	"eda-in-golang/baskets/internal/domain"
	"eda-in-golang/baskets/internal/grpc"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/rpc"
	"eda-in-golang/internal/web"
)

func TestProvider(t *testing.T) {
	var err error

	// init registry
	reg := registry.New()
	err = domain.Registrations(reg)
	if err != nil {
		t.Fatal(err)
	}
	// init repos
	baskets := domain.NewFakeBasketRepository()
	stores := domain.NewFakeStoreCacheRepository()
	products := domain.NewFakeProductCacheRepository()
	dispatcher := ddd.NewEventDispatcher[ddd.Event]()

	// init app
	app := application.New(baskets, stores, products, dispatcher)

	// start grpc
	rpcConfig := rpc.RpcConfig{
		Host: "0.0.0.0",
		Port: ":9095",
	}
	grpcServer := grpcstd.NewServer()
	// start rest
	webConfig := web.WebConfig{
		Host: "0.0.0.0",
		Port: ":9090",
	}
	mux := chi.NewMux()

	err = grpc.RegisterServer(app, grpcServer)
	if err != nil {
		t.Fatal(err)
	}
	err = RegisterGateway(context.Background(), mux, rpcConfig.Address())
	if err != nil {
		t.Fatal(err)
	}

	// start up the GRPC API; we proxy the REST api through the GRPC API
	{
		listener, err := net.Listen("tcp", rpcConfig.Address())
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			if err = grpcServer.Serve(listener); err != nil && err != grpcstd.ErrServerStopped {
				t.Error(err)
				return
			}
		}()
		defer func() {
			grpcServer.GracefulStop()
		}()
	}

	// start up the REST API
	{
		webServer := &http.Server{
			Addr:    webConfig.Address(),
			Handler: mux,
		}
		go func() {
			if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				t.Error(err)
				return
			}
		}()
		defer func() {
			if err := webServer.Shutdown(context.Background()); err != nil {
				t.Error(err)
				return
			}
		}()
	}

	pact := dsl.Pact{
		Provider: "baskets-api",
	}
	_, err = pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://%s", webConfig.Address()),
		ProviderVersion:            "1.0.0",
		BrokerURL:                  "http://127.0.0.1:9292",
		BrokerUsername:             "pactuser",
		BrokerPassword:             "pactpass",
		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
	})
	if err != nil {
		t.Error(err)
	}
}
