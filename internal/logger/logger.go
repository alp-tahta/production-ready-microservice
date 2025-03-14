package logger

import (
	"log/slog"
	"os"
)

func Init() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}
