package main

import (
	"github.com/artemKapitonov/avito_test_task/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		logrus.Fatalf("Can't start application: %s", err.Error())
	}
}
