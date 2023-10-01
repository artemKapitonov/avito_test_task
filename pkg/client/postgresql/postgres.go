package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/artemKapitonov/avito_test_task/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Client is an interface for PostgreSQL client operations
type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Ping(ctx context.Context) error
	Close()
}

// Config represents the configuration for connecting to PostgreSQL
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// ConnectionConfig creates a connection configuration for PostgreSQL
func ConnectionConfig(cfg Config) (*pgxpool.Config, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ConnectToDB connects to the PostgreSQL database
func ConnectToDB(ctx context.Context, cfg Config) (db *pgxpool.Pool, err error) {
	var maxAttempts = 5
	connCfg, err := ConnectionConfig(cfg)
	if err != nil {
		return nil, err
	}

	err = utils.DoWithTries(
		func() error {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			db, err = pgxpool.NewWithConfig(ctx, connCfg)
			if err != nil {
				return err
			}

			if err := db.Ping(ctx); err != nil {
				return err
			}

			return nil
		}, maxAttempts, 5*time.Second)

	if err != nil {
		logrus.Fatalf("Do with tries Error: %s", err.Error())
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	logrus.Println("Database connection successful")
	return db, nil
}
