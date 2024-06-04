package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"context"

	v1 "github.com/artemKapitonov/avito_test_task/internal/pkg/controller/http/v1"
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase"
	convert "github.com/artemKapitonov/avito_test_task/internal/pkg/usecase/currency_converter"
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase/storage"
	migrate "github.com/artemKapitonov/avito_test_task/migrations"
	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
	"github.com/artemKapitonov/avito_test_task/pkg/logging"
	httpserver "github.com/artemKapitonov/avito_test_task/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// App structure of application.
type App struct {
	log               *slog.Logger
	Controller        *v1.Controller
	UseCase           *usecase.UseCase
	Storage           *storage.Storage
	CurrencyConverter *convert.CurrencyConvert
	Server            *httpserver.Server
}

// New application.
func New() *App {
	const op = "app.New"

	var LoggerCfg = logging.Config{
		Level:   viper.GetString("log.level"),
		Handler: viper.GetString("log.handler"),
		Writer:  viper.GetString("log.writer"),
	}

	logger := logging.New(LoggerCfg)

	gin.SetMode(gin.TestMode)

	var app = &App{}

	app.log = logger.Logger

	log := app.log.With(slog.String("op", op))

	// Get API Layer token from environment variable.
	token := os.Getenv("API_LAYER_TOKEN")

	// Get port from configuration
	port := viper.GetString("port")

	// Create a new context.
	ctx := context.TODO()

	// Connect to the database.
	db, err := postgresql.ConnectToDB(ctx, postgresql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Error("Failed to connect to postgres database Error:", err)
	} else {
		log.Info("Database connection successful")
	}

	if err := migrate.Create(db); err != nil {
		log.Error("Can't create migrations Error:", err)
	}

	app.CurrencyConverter = convert.New(token, app.log)

	app.Storage = storage.New(db)

	app.UseCase = usecase.New(app.Storage, app.CurrencyConverter)

	app.Controller = v1.New(app.UseCase)

	app.Server = httpserver.New(app.Controller.InitRoutes(logger), port, app.log)

	return app
}

// Run is starting application.
func (a *App) Run() error {
	const op = "app.Run"

	log := a.log.With(slog.String("op", op))

	a.Server.Start()

	err := ShutdownApp(a)
	if err != nil {
		return err
	}

	log.Info("App Shutting down")

	return nil
}

// ShutdownApp is shutting down application.
func ShutdownApp(a *App) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := a.Server.Shutdown(context.Background()); err != nil {
		return err
	}

	defer a.Storage.Close()

	return nil
}
