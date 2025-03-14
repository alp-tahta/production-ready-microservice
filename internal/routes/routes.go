package routes

import (
	"net/http"

	"github.com/alp-tahta/production-ready-microservice/internal/handler"
)

func RegisterRoutes(mux *http.ServeMux, h *handler.Handler) {
	mux.HandleFunc("/health", h.Health)
}
