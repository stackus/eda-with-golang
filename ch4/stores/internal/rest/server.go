package rest

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
	"github.com/stackus/eda-with-golang/ch4/stores/internal/application"
	storespb "github.com/stackus/eda-with-golang/ch4/stores/internal/pb"
)

func Register(_ application.App, mono monolith.Monolith) error {
	// mount the stores GRPX gateway
	gateway := runtime.NewServeMux()
	storeAPIRoot := "/api/stores"
	mono.Mux().Mount(storeAPIRoot, gateway)
	err := storespb.RegisterStoresServiceHandlerFromEndpoint(mono.Context(), gateway, mono.Config().Rpc.Address(), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		return err
	}

	// mount the stores swagger UI
	storeUIRoot := "/stores-ui/"
	mono.Mux().Mount(storeUIRoot, http.StripPrefix(storeUIRoot, http.FileServer(http.FS(swaggerUI))))

	return nil
}
