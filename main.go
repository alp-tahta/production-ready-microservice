package main

import (
	"net/http"
	"os"

	"github.com/alp-tahta/production-ready-microservice/internal/handler"
	"github.com/alp-tahta/production-ready-microservice/internal/logger"
	"github.com/alp-tahta/production-ready-microservice/internal/routes"
	"github.com/alp-tahta/production-ready-microservice/internal/server"
)

func main() {
	port := ":8080"

	logger := logger.Init()

	mux := http.NewServeMux()
	handler := handler.New(logger)
	routes.RegisterRoutes(mux, handler)

	logger.Info("Starting HTTP server at", "port", port)
	err := server.Init(port, mux)
	if err != nil {
		logger.Error("Failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}
