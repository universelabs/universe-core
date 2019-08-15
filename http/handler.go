package http 

import (
	// stdlib

	// universe
	"github.com/universelabs/universe-core/universe"
	// deps
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// A collection of all the service handlers
type Handler struct {
	// router
	*chi.Mux
	// service handlers
	*KeyManagerHandler
}

// Instantiate the chi.Mux and mount the service handlers for chi's ServeHTTP
func NewHandler(km universe.KeyManager) {
	h := &Handler{}
	h.Mux = chi.NewRouter()
	h.Mux.Use(
		// CHANGE TO PROTOBUF
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	// instantiate and route handlers
	h.Route("/v", func(r chi.Router) {
		r.Mount("/api/keymanager", h.KeyManagerHandler)
	})

	return h
} 