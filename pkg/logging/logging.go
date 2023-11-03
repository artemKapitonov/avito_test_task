package logging

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	Logger *slog.Logger
	Writer io.Writer
}

func New() *Logger {
	err := os.RemoveAll("logs")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("logs", 0777)
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	return &Logger{
		Logger: slog.New(slog.NewJSONHandler(logFile, nil)),
		Writer: logFile,
	}
}
