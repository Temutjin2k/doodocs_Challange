package logger

import (
	"log"
	"log/slog"
	"os"
)

// Initialize the logger
func InitLogger() *slog.Logger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewTextHandler(file, nil))

	return logger
}
