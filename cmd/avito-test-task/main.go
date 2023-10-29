package main

import (
	"log/slog"

	"github.com/artemKapitonov/avito_test_task/internal/app"
)

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		slog.Error("Can't start application Error:", err)
	}
}
