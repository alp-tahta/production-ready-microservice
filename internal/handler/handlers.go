package handler

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	l *slog.Logger
}

func New(l *slog.Logger) *Handler {
	return &Handler{
		l: l,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		h.l.Error("Failed to write response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
