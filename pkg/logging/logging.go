package logging

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	if err := os.MkdirAll("logs", 0777); err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	return logger
}
