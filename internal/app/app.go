package app

import (
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
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// App structure of application.
type App struct {
	Controller        *v1.Controller
	UseCase           *usecase.UseCase
	Storage           *storage.Storage
	CurrencyConverter *convert.CurrencyConvert
	Server            *httpserver.Server
}

// New application.
func New() *App {
	// Set Gin mode to TestMode.
	gin.SetMode(gin.TestMode)

	// Set logrus formatter to JSONFormatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Initialize configurations
	if err := config.Init(); err != nil {
		logrus.Fatalf("Can't init configs Error: %s", err.Error())
	}

	// Load environment variables from .env file.
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("Can't load env Error: %s", err.Error())
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
		logrus.Fatalf("Can't connect to database Error: %s", err.Error())
	}

	if err := migrate.Create(db); err != nil {
		logrus.Fatalf("Can't create migrations Error: %s", err.Error())
	}

	app.CurrencyConverter = convert.New(token)

	app.Storage = storage.New(db)

	app.UseCase = usecase.New(app.Storage, app.CurrencyConverter)

	app.Controller = v1.New(app.UseCase)

	app.Server = httpserver.New(app.Controller.InitRoutes(), port)

	return app
}

// Run application.
func (a *App) Run() error {
	if err := a.Server.Start(); err != nil {
		return err
	}

	err := ShutdownApp(a)
	if err != nil {
		return err
	}

	return nil
}

// ShutdownApp shutting down application.
func ShutdownApp(a *App) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := a.Server.Shutdown(context.Background()); err != nil {
		return err
	}

	defer a.Storage.Close()

	logrus.Println("App Shuting down")

	return nil
}
