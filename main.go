package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/alp-tahta/production-ready-microservice/internal/handler"
	"github.com/alp-tahta/production-ready-microservice/internal/routes"
	"github.com/alp-tahta/production-ready-microservice/internal/server"
)

func main() {
	port := ":8080"
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Starting HTTP server at", "port", port)

	mux := http.NewServeMux()
	handler := handler.New(logger)
	routes.RegisterRoutes(mux, handler)

	err := server.Init(port, mux)
	if err != nil {
		logger.Error("Failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}
