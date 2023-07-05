package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/artemKapitonov/avito_test_task/internal/config"
	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := config.Init(); err != nil {
		logrus.Fatalf("Can't init configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
	}

	cfg := postgresql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		logrus.Fatalf("Can't connection with database: %s", err.Error())
	}

	defer func() {
		if err := db.Close(); err != nil {
			logrus.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := db.Ping(); err != nil {
		logrus.Fatalf(err.Error())
	}
	logrus.Println("DB connection success")

	if err := goose.SetDialect("postgres"); err != nil {
		logrus.Fatal("Error in set dialect")
	}

	// выполнение миграций
	if err := goose.Up(db, "migrations/schema"); err != nil {
		logrus.Fatal("Ошибка при выполнении миграций:", err)
	}
}
