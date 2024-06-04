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
		slog.Error("Can't start application Error:", err)
	}
}

func InitConfig() {
	// Initialize configurations
	if err := config.Init(); err != nil {
		panic(err)
	}

	// Load environment variables from .env file.
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
