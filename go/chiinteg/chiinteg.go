package chiinteg

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaz/pprotein/integration"
)

// Integrate sets up debugging for the chi router
func Integrate(r *chi.Mux) {
	EnableDebugHandler(r)
	EnableDebugMode(r)
}

// EnableDebugHandler registers debug handlers to the router
func EnableDebugHandler(r *chi.Mux) {
	integration.RegisterDebugHandlers(r)
}

// EnableDebugMode enables debugging mode (currently a no-op)
func EnableDebugMode(r *chi.Mux) {
	// This function does nothing for now
	return
}
