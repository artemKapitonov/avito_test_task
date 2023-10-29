package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"context"

	"github.com/artemKapitonov/avito_test_task/internal/config"
	v1 "github.com/artemKapitonov/avito_test_task/internal/pkg/controller/http/v1"
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase"
	convert "github.com/artemKapitonov/avito_test_task/internal/pkg/usecase/currency_converter"
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase/storage"
	migrate "github.com/artemKapitonov/avito_test_task/migrations"
	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
	httpserver "github.com/artemKapitonov/avito_test_task/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// App structure of application.
type App struct {
	logger            *slog.Logger
	Controller        *v1.Controller
	UseCase           *usecase.UseCase
	Storage           *storage.Storage
	CurrencyConverter *convert.CurrencyConvert
	Server            *httpserver.Server
}

// New application.
func New() *App {

	file, err := os.Open("logs/all.log")
	if err != nil {
		return nil
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// Set Gin mode to TestMode.
	gin.SetMode(gin.TestMode)

	slog.SetDefault(logger)

	// Initialize configurations
	if err := config.Init(); err != nil {
		logger.Error("Can't init configs Error: %s", err.Error())
	}

	// Load environment variables from .env file.
	if err := godotenv.Load(".env"); err != nil {
		logger.Error("Can't load env Error: %s", err.Error())
	}

	app := &App{}

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
		logger.Error("Can't connect to database Error: %s", err.Error())
	}

	logger.Info("Database connection successful")

	if err := migrate.Create(db); err != nil {
		logger.Error("Can't create migrations Error: %s", err.Error())
	}

	app.logger = logger

	app.CurrencyConverter = convert.New(token, app.logger)

	app.Storage = storage.New(db)

	app.UseCase = usecase.New(app.Storage, app.CurrencyConverter)

	app.Controller = v1.New(app.UseCase)

	app.Server = httpserver.New(app.Controller.InitRoutes(), port, app.logger)

	return app
}

// Run is starting application.
func (a *App) Run() error {
	if err := a.Server.Start(); err != nil {
		return err
	}

	err := ShutdownApp(a)
	if err != nil {
		return err
	}
	a.logger.Info("App Shutting down")

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
