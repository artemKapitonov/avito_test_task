package main

import (
	"log/slog"

	"github.com/artemKapitonov/avito_test_task/internal/app"
	"github.com/artemKapitonov/avito_test_task/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	InitConfig()

	a := app.New()

	if err := a.Run(); err != nil {
		slog.Error("Can't start application", "error", err)
	}
}

func InitConfig() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
