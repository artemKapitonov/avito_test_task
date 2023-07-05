package app

import (
	"os"

	"context"

	"github.com/artemKapitonov/avito_test_task/internal/config"
	v1 "github.com/artemKapitonov/avito_test_task/internal/pkg/controller/http/v1"
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase"
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase/storage"
	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
	httpserver "github.com/artemKapitonov/avito_test_task/pkg/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	Controller *v1.Controller
	UseCase    *usecase.UseCase
	Storage    *storage.Storage
	Server     *httpserver.Server
}

func New() *App {
	if err := config.Init(); err != nil {
		logrus.Fatalf("Can't init configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
	}

	port := viper.GetString("port")

	app := &App{}

	db, err := postgresql.ConnectToDB(context.TODO(), postgresql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Can't connect to database: %s", err.Error())
	}

	app.Storage = storage.New(db)

	app.UseCase = usecase.New(app.Storage)

	app.Controller = v1.New(app.UseCase)

	app.Server = httpserver.New(app.Controller.InitRoutes(), port)

	return app
}

func (a *App) Run() error {
	return a.Server.Start()
}
