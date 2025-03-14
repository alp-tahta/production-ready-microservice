package handler

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	l *slog.Logger
}

func New(l *slog.Logger) *Handler {
	return &Handler{l: l}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Health check log test") // TODO remove
	w.Write([]byte("OK"))
}
