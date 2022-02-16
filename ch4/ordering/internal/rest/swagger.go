package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterSwagger(mux *chi.Mux) error {
	const specRoot = "/ordering-spec/"

	// mount the swagger specification
	mux.Mount(specRoot, http.StripPrefix(specRoot, http.FileServer(http.FS(swaggerUI))))

	return nil
}
