package main

import (
	"github.com/artemKapitonov/avito_test_task/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	app := app.New()

	if err := app.Run(); err != nil {
		logrus.Fatalf("Can't start application: %s", err.Error())
	}
}
 